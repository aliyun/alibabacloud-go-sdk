// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDpiConfigErrorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDpiConfigError(v []*ListDpiConfigErrorResponseBodyDpiConfigError) *ListDpiConfigErrorResponseBody
	GetDpiConfigError() []*ListDpiConfigErrorResponseBodyDpiConfigError
	SetMaxResults(v int32) *ListDpiConfigErrorResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDpiConfigErrorResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDpiConfigErrorResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListDpiConfigErrorResponseBody
	GetTotal() *int32
}

type ListDpiConfigErrorResponseBody struct {
	// A list of DPI configuration errors.
	DpiConfigError []*ListDpiConfigErrorResponseBodyDpiConfigError `json:"DpiConfigError,omitempty" xml:"DpiConfigError,omitempty" type:"Repeated"`
	// The maximum number of configuration errors to return on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F47B5293-27B6-48EF-A9C6-E90A41449813
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DPI configuration errors.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListDpiConfigErrorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDpiConfigErrorResponseBody) GoString() string {
	return s.String()
}

func (s *ListDpiConfigErrorResponseBody) GetDpiConfigError() []*ListDpiConfigErrorResponseBodyDpiConfigError {
	return s.DpiConfigError
}

func (s *ListDpiConfigErrorResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDpiConfigErrorResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDpiConfigErrorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDpiConfigErrorResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDpiConfigErrorResponseBody) SetDpiConfigError(v []*ListDpiConfigErrorResponseBodyDpiConfigError) *ListDpiConfigErrorResponseBody {
	s.DpiConfigError = v
	return s
}

func (s *ListDpiConfigErrorResponseBody) SetMaxResults(v int32) *ListDpiConfigErrorResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDpiConfigErrorResponseBody) SetNextToken(v string) *ListDpiConfigErrorResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDpiConfigErrorResponseBody) SetRequestId(v string) *ListDpiConfigErrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDpiConfigErrorResponseBody) SetTotal(v int32) *ListDpiConfigErrorResponseBody {
	s.Total = &v
	return s
}

func (s *ListDpiConfigErrorResponseBody) Validate() error {
	if s.DpiConfigError != nil {
		for _, item := range s.DpiConfigError {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDpiConfigErrorResponseBodyDpiConfigError struct {
	// The type of the configuration error.
	//
	// - DeviceNotSupported: The Smart Access Gateway device does not support the DPI feature.
	//
	// - VersionNotSupported: The DPI version of the Smart Access Gateway device is too old.
	//
	// - **NotEnable**: The DPI feature is disabled for the Smart Access Gateway device.
	//
	// example:
	//
	// DeviceNotSupported
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// A list of rule configuration errors.
	RuleConfigErrorList []*ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList `json:"RuleConfigErrorList,omitempty" xml:"RuleConfigErrorList,omitempty" type:"Repeated"`
	// The serial number of the Smart Access Gateway device.
	//
	// example:
	//
	// sag-2160808****
	SN *string `json:"SN,omitempty" xml:"SN,omitempty"`
	// The ID of the Smart Access Gateway instance.
	//
	// example:
	//
	// sag-1e8sgws6b133b8****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ListDpiConfigErrorResponseBodyDpiConfigError) String() string {
	return dara.Prettify(s)
}

func (s ListDpiConfigErrorResponseBodyDpiConfigError) GoString() string {
	return s.String()
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) GetErrorType() *string {
	return s.ErrorType
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) GetRuleConfigErrorList() []*ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList {
	return s.RuleConfigErrorList
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) GetSN() *string {
	return s.SN
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) SetErrorType(v string) *ListDpiConfigErrorResponseBodyDpiConfigError {
	s.ErrorType = &v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) SetRuleConfigErrorList(v []*ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) *ListDpiConfigErrorResponseBodyDpiConfigError {
	s.RuleConfigErrorList = v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) SetSN(v string) *ListDpiConfigErrorResponseBodyDpiConfigError {
	s.SN = &v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) SetSmartAGId(v string) *ListDpiConfigErrorResponseBodyDpiConfigError {
	s.SmartAGId = &v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigError) Validate() error {
	if s.RuleConfigErrorList != nil {
		for _, item := range s.RuleConfigErrorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList struct {
	// A list of IDs of application groups that have configuration errors.
	DpiGroupIds []*string `json:"DpiGroupIds,omitempty" xml:"DpiGroupIds,omitempty" type:"Repeated"`
	// A list of IDs of applications that have configuration errors.
	DpiSignatureIds []*string `json:"DpiSignatureIds,omitempty" xml:"DpiSignatureIds,omitempty" type:"Repeated"`
	// The ID of the rule that is associated with the application that has a configuration error.
	//
	// - If you query DPI configuration errors for Resource Access Management, this parameter indicates the ID of the Resource Access Management rule instance that has a configuration error.
	//
	// - If you query DPI configuration errors for a QoS policy, this parameter indicates the ID of the quintuple rule instance that has a configuration error.
	//
	// example:
	//
	// qospy-axud4s62gz632b****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) String() string {
	return dara.Prettify(s)
}

func (s ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) GoString() string {
	return s.String()
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) GetDpiGroupIds() []*string {
	return s.DpiGroupIds
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) GetDpiSignatureIds() []*string {
	return s.DpiSignatureIds
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) GetRuleId() *string {
	return s.RuleId
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) SetDpiGroupIds(v []*string) *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList {
	s.DpiGroupIds = v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) SetDpiSignatureIds(v []*string) *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList {
	s.DpiSignatureIds = v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) SetRuleId(v string) *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList {
	s.RuleId = &v
	return s
}

func (s *ListDpiConfigErrorResponseBodyDpiConfigErrorRuleConfigErrorList) Validate() error {
	return dara.Validate(s)
}
