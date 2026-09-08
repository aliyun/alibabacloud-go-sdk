// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWebhookContactResponseBody
	GetCode() *string
	SetMessage(v string) *CreateWebhookContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWebhookContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWebhookContactResponseBody
	GetSuccess() *bool
}

type CreateWebhookContactResponseBody struct {
	// The error code returned if the call failed. For more information, see error codes.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned if the call failed.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A5F62766-1C2F-1F56-A39D-63E3D30F0633
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWebhookContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWebhookContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWebhookContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWebhookContactResponseBody) SetCode(v string) *CreateWebhookContactResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWebhookContactResponseBody) SetMessage(v string) *CreateWebhookContactResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWebhookContactResponseBody) SetRequestId(v string) *CreateWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWebhookContactResponseBody) SetSuccess(v bool) *CreateWebhookContactResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWebhookContactResponseBody) Validate() error {
	return dara.Validate(s)
}
