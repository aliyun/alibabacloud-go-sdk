// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *SubmitCampaignResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *SubmitCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitCampaignResponseBody
	GetRequestId() *string
}

type SubmitCampaignResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7CC6523B-0E51-1B62-8DA5-6A9831CAE315
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitCampaignResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *SubmitCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCampaignResponseBody) SetCode(v string) *SubmitCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetHttpStatusCode(v string) *SubmitCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetMessage(v string) *SubmitCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitCampaignResponseBody) SetRequestId(v string) *SubmitCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
