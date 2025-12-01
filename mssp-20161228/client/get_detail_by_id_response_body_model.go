// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetailByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDetailByIdResponseBody
	GetCode() *string
	SetData(v *GetDetailByIdResponseBodyData) *GetDetailByIdResponseBody
	GetData() *GetDetailByIdResponseBodyData
	SetHttpStatusCode(v int32) *GetDetailByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDetailByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDetailByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDetailByIdResponseBody
	GetSuccess() *bool
}

type GetDetailByIdResponseBody struct {
	// Interface return code.
	//
	// example:
	//
	// 404
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data query result.
	Data *GetDetailByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DAB46EC5-3746-59C4-B6D2-469F442EC73F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values: - **true**: indicates a successful call. - **false**: indicates a failed call.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDetailByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDetailByIdResponseBody) GetData() *GetDetailByIdResponseBodyData {
	return s.Data
}

func (s *GetDetailByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDetailByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDetailByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDetailByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDetailByIdResponseBody) SetCode(v string) *GetDetailByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetData(v *GetDetailByIdResponseBodyData) *GetDetailByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetDetailByIdResponseBody) SetHttpStatusCode(v int32) *GetDetailByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetMessage(v string) *GetDetailByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetRequestId(v string) *GetDetailByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDetailByIdResponseBody) SetSuccess(v bool) *GetDetailByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetDetailByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDetailByIdResponseBodyData struct {
	// Vulnerability details.
	VulDetails []*GetDetailByIdResponseBodyDataVulDetails `json:"VulDetails,omitempty" xml:"VulDetails,omitempty" type:"Repeated"`
}

func (s GetDetailByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyData) GetVulDetails() []*GetDetailByIdResponseBodyDataVulDetails {
	return s.VulDetails
}

func (s *GetDetailByIdResponseBodyData) SetVulDetails(v []*GetDetailByIdResponseBodyDataVulDetails) *GetDetailByIdResponseBodyData {
	s.VulDetails = v
	return s
}

func (s *GetDetailByIdResponseBodyData) Validate() error {
	if s.VulDetails != nil {
		for _, item := range s.VulDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDetailByIdResponseBodyDataVulDetails struct {
	// CVE ID.
	//
	// example:
	//
	// CVE-2022-21291
	CveId *string `json:"CveId,omitempty" xml:"CveId,omitempty"`
	// The CVSS score of the vulnerability in the Alibaba Cloud vulnerability database.
	//
	// example:
	//
	// 10.0
	CvssScore *string `json:"CvssScore,omitempty" xml:"CvssScore,omitempty"`
	// Fix suggestion.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/CVE-2022-21291
	FixSuggestion *string `json:"FixSuggestion,omitempty" xml:"FixSuggestion,omitempty"`
	// Title of the vulnerability announcement.
	//
	// example:
	//
	// Chanjet T-Plus SetupAccount/Upload. Aspx file upload vulnerability(CNVD-2022-60632)
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetDetailByIdResponseBodyDataVulDetails) String() string {
	return dara.Prettify(s)
}

func (s GetDetailByIdResponseBodyDataVulDetails) GoString() string {
	return s.String()
}

func (s *GetDetailByIdResponseBodyDataVulDetails) GetCveId() *string {
	return s.CveId
}

func (s *GetDetailByIdResponseBodyDataVulDetails) GetCvssScore() *string {
	return s.CvssScore
}

func (s *GetDetailByIdResponseBodyDataVulDetails) GetFixSuggestion() *string {
	return s.FixSuggestion
}

func (s *GetDetailByIdResponseBodyDataVulDetails) GetTitle() *string {
	return s.Title
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCveId(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CveId = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetCvssScore(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.CvssScore = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetFixSuggestion(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.FixSuggestion = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) SetTitle(v string) *GetDetailByIdResponseBodyDataVulDetails {
	s.Title = &v
	return s
}

func (s *GetDetailByIdResponseBodyDataVulDetails) Validate() error {
	return dara.Validate(s)
}
