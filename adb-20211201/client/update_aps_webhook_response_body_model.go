// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApsWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateApsWebhookResponseBody
	GetCode() *string
	SetDBClusterId(v string) *UpdateApsWebhookResponseBody
	GetDBClusterId() *string
	SetData(v string) *UpdateApsWebhookResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateApsWebhookResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateApsWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateApsWebhookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateApsWebhookResponseBody
	GetSuccess() *bool
}

type UpdateApsWebhookResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// exampleRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateApsWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateApsWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateApsWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateApsWebhookResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateApsWebhookResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateApsWebhookResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateApsWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateApsWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateApsWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateApsWebhookResponseBody) SetCode(v string) *UpdateApsWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetDBClusterId(v string) *UpdateApsWebhookResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetData(v string) *UpdateApsWebhookResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetHttpStatusCode(v int32) *UpdateApsWebhookResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetMessage(v string) *UpdateApsWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetRequestId(v string) *UpdateApsWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) SetSuccess(v bool) *UpdateApsWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateApsWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
