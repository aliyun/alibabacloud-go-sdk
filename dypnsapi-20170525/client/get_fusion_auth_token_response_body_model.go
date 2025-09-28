// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFusionAuthTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFusionAuthTokenResponseBody
	GetCode() *string
	SetMessage(v string) *GetFusionAuthTokenResponseBody
	GetMessage() *string
	SetModel(v string) *GetFusionAuthTokenResponseBody
	GetModel() *string
	SetRequestId(v string) *GetFusionAuthTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFusionAuthTokenResponseBody
	GetSuccess() *bool
}

type GetFusionAuthTokenResponseBody struct {
	// The response code. If OK is returned, the request is successful. Other values indicate that the request failed. For more information, see Error codes.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The authentication code. The value of this parameter is a string.
	//
	// example:
	//
	// FKcksloqk***********jalEc+
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CC3BB6D2-2FDF-4321-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFusionAuthTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFusionAuthTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFusionAuthTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFusionAuthTokenResponseBody) GetModel() *string {
	return s.Model
}

func (s *GetFusionAuthTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFusionAuthTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFusionAuthTokenResponseBody) SetCode(v string) *GetFusionAuthTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetMessage(v string) *GetFusionAuthTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetModel(v string) *GetFusionAuthTokenResponseBody {
	s.Model = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetRequestId(v string) *GetFusionAuthTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) SetSuccess(v bool) *GetFusionAuthTokenResponseBody {
	s.Success = &v
	return s
}

func (s *GetFusionAuthTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
