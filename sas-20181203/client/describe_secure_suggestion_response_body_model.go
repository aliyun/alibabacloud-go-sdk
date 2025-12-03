// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecureSuggestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalTime(v int64) *DescribeSecureSuggestionResponseBody
	GetCalTime() *int64
	SetRequestId(v string) *DescribeSecureSuggestionResponseBody
	GetRequestId() *string
	SetSuggestions(v []*DescribeSecureSuggestionResponseBodySuggestions) *DescribeSecureSuggestionResponseBody
	GetSuggestions() []*DescribeSecureSuggestionResponseBodySuggestions
	SetTotalCount(v int32) *DescribeSecureSuggestionResponseBody
	GetTotalCount() *int32
}

type DescribeSecureSuggestionResponseBody struct {
	// example:
	//
	// 1755744253000
	CalTime *int64 `json:"CalTime,omitempty" xml:"CalTime,omitempty"`
	// example:
	//
	// 676F80E3-4B3F-43DA-9CBB-5FF79F202AA2
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Suggestions []*DescribeSecureSuggestionResponseBodySuggestions `json:"Suggestions,omitempty" xml:"Suggestions,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecureSuggestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBody) GetCalTime() *int64 {
	return s.CalTime
}

func (s *DescribeSecureSuggestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecureSuggestionResponseBody) GetSuggestions() []*DescribeSecureSuggestionResponseBodySuggestions {
	return s.Suggestions
}

func (s *DescribeSecureSuggestionResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecureSuggestionResponseBody) SetCalTime(v int64) *DescribeSecureSuggestionResponseBody {
	s.CalTime = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetRequestId(v string) *DescribeSecureSuggestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetSuggestions(v []*DescribeSecureSuggestionResponseBodySuggestions) *DescribeSecureSuggestionResponseBody {
	s.Suggestions = v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) SetTotalCount(v int32) *DescribeSecureSuggestionResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBody) Validate() error {
	if s.Suggestions != nil {
		for _, item := range s.Suggestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecureSuggestionResponseBodySuggestions struct {
	Detail []*DescribeSecureSuggestionResponseBodySuggestionsDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// example:
	//
	// 40
	Points *int32 `json:"Points,omitempty" xml:"Points,omitempty"`
	// example:
	//
	// SS_ALARM
	SuggestType *string `json:"SuggestType,omitempty" xml:"SuggestType,omitempty"`
}

func (s DescribeSecureSuggestionResponseBodySuggestions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestions) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetDetail() []*DescribeSecureSuggestionResponseBodySuggestionsDetail {
	return s.Detail
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetPoints() *int32 {
	return s.Points
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) GetSuggestType() *string {
	return s.SuggestType
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetDetail(v []*DescribeSecureSuggestionResponseBodySuggestionsDetail) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Detail = v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetPoints(v int32) *DescribeSecureSuggestionResponseBodySuggestions {
	s.Points = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) SetSuggestType(v string) *DescribeSecureSuggestionResponseBodySuggestions {
	s.SuggestType = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestions) Validate() error {
	if s.Detail != nil {
		for _, item := range s.Detail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecureSuggestionResponseBodySuggestionsDetail struct {
	// example:
	//
	// Malicious tampering of Web pages will affect your normal access to web page content, and may also lead to serious economic losses, brand losses, and even political risks. The webpage tamper-proof service can monitor the website directory in real time and restore the tampered files or directories through backup, so as to ensure that the website information of important systems is not tampered with maliciously and prevent the occurrence of horse hanging, black chain, illegal implantation of terrorist threats, pornography and other content.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// REINFORCE_WEB_LOCK
	SubType *string `json:"SubType,omitempty" xml:"SubType,omitempty"`
	// example:
	//
	// Website tamper-proofing capability not configured
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponseBodySuggestionsDetail) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetSubType() *string {
	return s.SubType
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) GetTitle() *string {
	return s.Title
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetDescription(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Description = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetSubType(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.SubType = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) SetTitle(v string) *DescribeSecureSuggestionResponseBodySuggestionsDetail {
	s.Title = &v
	return s
}

func (s *DescribeSecureSuggestionResponseBodySuggestionsDetail) Validate() error {
	return dara.Validate(s)
}
