// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *UpdateCampaignResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *UpdateCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCampaignResponseBody
	GetRequestId() *string
}

type UpdateCampaignResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// BC976D32-AC4C-4E0F-8AA9-F4BC6C4E2B3E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateCampaignResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *UpdateCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCampaignResponseBody) SetCode(v string) *UpdateCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCampaignResponseBody) SetHttpStatusCode(v int64) *UpdateCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateCampaignResponseBody) SetMessage(v string) *UpdateCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCampaignResponseBody) SetRequestId(v string) *UpdateCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
