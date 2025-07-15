// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResumeCampaignResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *ResumeCampaignResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ResumeCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResumeCampaignResponseBody
	GetRequestId() *string
}

type ResumeCampaignResponseBody struct {
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
	// F505C4C8-1E12-573A-9BA7-4BEAAD129553
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResumeCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResumeCampaignResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ResumeCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResumeCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeCampaignResponseBody) SetCode(v string) *ResumeCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetHttpStatusCode(v string) *ResumeCampaignResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetMessage(v string) *ResumeCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeCampaignResponseBody) SetRequestId(v string) *ResumeCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeCampaignResponseBody) Validate() error {
	return dara.Validate(s)
}
