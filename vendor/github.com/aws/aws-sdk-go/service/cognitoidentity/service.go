// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cognitoidentity

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon Cognito is a web service that delivers scoped temporary credentials
// to mobile devices and other untrusted environments. Amazon Cognito uniquely
// identifies a device and supplies the user with a consistent identity over
// the lifetime of an application.
//
// Using Amazon Cognito, you can enable authentication with one or more third-party
// identity providers (Facebook, Google, or Login with Amazon), and you can
// also choose to support unauthenticated access from your app. Cognito delivers
// a unique identifier for each user and acts as an OpenID token provider trusted
// by AWS Security Token Service (STS) to access temporary, limited-privilege
// AWS credentials.
//
// To provide end-user credentials, first make an unsigned call to GetId. If
// the end user is authenticated with one of the supported identity providers,
// set the Logins map with the identity provider token. GetId returns a unique
// identifier for the user.
//
// Next, make an unsigned call to GetCredentialsForIdentity. This call expects
// the same Logins map as the GetId call, as well as the IdentityID originally
// returned by GetId. Assuming your identity pool has been configured via the
// SetIdentityPoolRoles operation, GetCredentialsForIdentity will return AWS
// credentials for your use. If your pool has not been configured with SetIdentityPoolRoles,
// or if you want to follow legacy flow, make an unsigned call to GetOpenIdToken,
// which returns the OpenID token necessary to call STS and retrieve AWS credentials.
// This call expects the same Logins map as the GetId call, as well as the IdentityID
// originally returned by GetId. The token returned by GetOpenIdToken can be
// passed to the STS operation AssumeRoleWithWebIdentity (http://docs.aws.amazon.com/STS/latest/APIReference/API_AssumeRoleWithWebIdentity.html)
// to retrieve AWS credentials.
//
// If you want to use Amazon Cognito in an Android, iOS, or Unity application,
// you will probably want to make API calls via the AWS Mobile SDK. To learn
// more, see the AWS Mobile SDK Developer Guide (http://docs.aws.amazon.com/mobile/index.html).
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type CognitoIdentity struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "cognito-identity"

// New creates a new instance of the CognitoIdentity client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a CognitoIdentity client from just a session.
//     svc := cognitoidentity.New(mySession)
//
//     // Create a CognitoIdentity client with additional configuration
//     svc := cognitoidentity.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *CognitoIdentity {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *CognitoIdentity {
	svc := &CognitoIdentity{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2014-06-30",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSCognitoIdentityService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBack(v4.Sign)
	svc.Handlers.Build.PushBack(jsonrpc.Build)
	svc.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	svc.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	svc.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a CognitoIdentity operation and runs any
// custom request initialization.
func (c *CognitoIdentity) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
