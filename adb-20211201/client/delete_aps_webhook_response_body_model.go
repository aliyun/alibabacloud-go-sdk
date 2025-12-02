// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApsWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApsWebhookResponseBody
	GetCode() *string
	SetDBClusterId(v string) *DeleteApsWebhookResponseBody
	GetDBClusterId() *string
	SetData(v string) *DeleteApsWebhookResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteApsWebhookResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteApsWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApsWebhookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApsWebhookResponseBody
	GetSuccess() *bool
}

type DeleteApsWebhookResponseBody struct {
	// API status or POP error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-uf63i4ij56b***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The returned data.
	//
	// example:
	//
	// []
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// exampleRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApsWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApsWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApsWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApsWebhookResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteApsWebhookResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteApsWebhookResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteApsWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApsWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApsWebhookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApsWebhookResponseBody) SetCode(v string) *DeleteApsWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetDBClusterId(v string) *DeleteApsWebhookResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetData(v string) *DeleteApsWebhookResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetHttpStatusCode(v int32) *DeleteApsWebhookResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetMessage(v string) *DeleteApsWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetRequestId(v string) *DeleteApsWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) SetSuccess(v bool) *DeleteApsWebhookResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApsWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
