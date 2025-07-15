// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AbortCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *AbortCampaignResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *AbortCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *AbortCampaignResponseBody
	GetRequestId() *string
}

type AbortCampaignResponseBody struct {
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

func (s AbortCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AbortCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *AbortCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *AbortCampaignResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *AbortCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AbortCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AbortCampaignResponseBody) SetCode(v string) *AbortCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *AbortCampaignResponseBody) SetHttpStatusCode(v string) *AbortCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AbortCampaignResponseBody) SetMessage(v string) *AbortCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *AbortCampaignResponseBody) SetRequestId(v string) *AbortCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
