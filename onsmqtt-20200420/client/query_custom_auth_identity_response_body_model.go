// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthIdentityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryCustomAuthIdentityResponseBody
	GetCode() *int32
	SetData(v *QueryCustomAuthIdentityResponseBodyData) *QueryCustomAuthIdentityResponseBody
	GetData() *QueryCustomAuthIdentityResponseBodyData
	SetMessage(v string) *QueryCustomAuthIdentityResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomAuthIdentityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomAuthIdentityResponseBody
	GetSuccess() *bool
}

type QueryCustomAuthIdentityResponseBody struct {
	// The response code. 200: The request is successful. Other codes indicate errors. For more information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *QueryCustomAuthIdentityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 11568B5B-13A8-4E72-9DBA-3A14F7D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. true: The request is successful. false: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomAuthIdentityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthIdentityResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthIdentityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryCustomAuthIdentityResponseBody) GetData() *QueryCustomAuthIdentityResponseBodyData {
	return s.Data
}

func (s *QueryCustomAuthIdentityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomAuthIdentityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomAuthIdentityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomAuthIdentityResponseBody) SetCode(v int32) *QueryCustomAuthIdentityResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBody) SetData(v *QueryCustomAuthIdentityResponseBodyData) *QueryCustomAuthIdentityResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomAuthIdentityResponseBody) SetMessage(v string) *QueryCustomAuthIdentityResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBody) SetRequestId(v string) *QueryCustomAuthIdentityResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBody) SetSuccess(v bool) *QueryCustomAuthIdentityResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomAuthIdentityResponseBodyData struct {
	// If more results are available, a NextToken is returned.
	//
	// example:
	//
	// AAAAAXA+GzVqTutYpgkFjBrchKzuvSbpuTqtt6OF9tsC9QnJ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The returned data.
	Results []*QueryCustomAuthIdentityResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s QueryCustomAuthIdentityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthIdentityResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthIdentityResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthIdentityResponseBodyData) GetResults() []*QueryCustomAuthIdentityResponseBodyDataResults {
	return s.Results
}

func (s *QueryCustomAuthIdentityResponseBodyData) SetNextToken(v string) *QueryCustomAuthIdentityResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyData) SetResults(v []*QueryCustomAuthIdentityResponseBodyDataResults) *QueryCustomAuthIdentityResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyData) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCustomAuthIdentityResponseBodyDataResults struct {
	// The client ID of the device when the identity type is CLIENT.
	//
	// example:
	//
	// GID_ICP@@@4d378084
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The identity type.
	//
	// - USER: Username only.
	//
	// - CLIENT: Specific to a client.
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The secret key.
	//
	// example:
	//
	// 62a5916d71767185b48907d85c2efa****
	Secret *string `json:"Secret,omitempty" xml:"Secret,omitempty"`
	// The signature verification mode. ORIGIN: Original value comparison, which compares the password with the secret. SIGNED: The HmacSHA1 signature is calculated for the clientId using the secret for verification, and then the password is compared.
	//
	// example:
	//
	// SIGNED
	SignMode *string `json:"SignMode,omitempty" xml:"SignMode,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s QueryCustomAuthIdentityResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthIdentityResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) GetClientId() *string {
	return s.ClientId
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) GetIdentityType() *string {
	return s.IdentityType
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) GetSecret() *string {
	return s.Secret
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) GetSignMode() *string {
	return s.SignMode
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) GetUsername() *string {
	return s.Username
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) SetClientId(v string) *QueryCustomAuthIdentityResponseBodyDataResults {
	s.ClientId = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) SetIdentityType(v string) *QueryCustomAuthIdentityResponseBodyDataResults {
	s.IdentityType = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) SetSecret(v string) *QueryCustomAuthIdentityResponseBodyDataResults {
	s.Secret = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) SetSignMode(v string) *QueryCustomAuthIdentityResponseBodyDataResults {
	s.SignMode = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) SetUsername(v string) *QueryCustomAuthIdentityResponseBodyDataResults {
	s.Username = &v
	return s
}

func (s *QueryCustomAuthIdentityResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
