// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMobilesCardSupportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CheckMobilesCardSupportResponseBody
	GetCode() *string
	SetData(v *CheckMobilesCardSupportResponseBodyData) *CheckMobilesCardSupportResponseBody
	GetData() *CheckMobilesCardSupportResponseBodyData
	SetRequestId(v string) *CheckMobilesCardSupportResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckMobilesCardSupportResponseBody
	GetSuccess() *bool
}

type CheckMobilesCardSupportResponseBody struct {
	// The HTTP status code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/101346.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CheckMobilesCardSupportResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 819BE656-D2E0-4858-8B21-B2E477085AAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckMobilesCardSupportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBody) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckMobilesCardSupportResponseBody) GetData() *CheckMobilesCardSupportResponseBodyData {
	return s.Data
}

func (s *CheckMobilesCardSupportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckMobilesCardSupportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckMobilesCardSupportResponseBody) SetCode(v string) *CheckMobilesCardSupportResponseBody {
	s.Code = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetData(v *CheckMobilesCardSupportResponseBodyData) *CheckMobilesCardSupportResponseBody {
	s.Data = v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetRequestId(v string) *CheckMobilesCardSupportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) SetSuccess(v bool) *CheckMobilesCardSupportResponseBody {
	s.Success = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckMobilesCardSupportResponseBodyData struct {
	// The list of returned results.
	QueryResult []*CheckMobilesCardSupportResponseBodyDataQueryResult `json:"queryResult,omitempty" xml:"queryResult,omitempty" type:"Repeated"`
}

func (s CheckMobilesCardSupportResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBodyData) GetQueryResult() []*CheckMobilesCardSupportResponseBodyDataQueryResult {
	return s.QueryResult
}

func (s *CheckMobilesCardSupportResponseBodyData) SetQueryResult(v []*CheckMobilesCardSupportResponseBodyDataQueryResult) *CheckMobilesCardSupportResponseBodyData {
	s.QueryResult = v
	return s
}

func (s *CheckMobilesCardSupportResponseBodyData) Validate() error {
	if s.QueryResult != nil {
		for _, item := range s.QueryResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckMobilesCardSupportResponseBodyDataQueryResult struct {
	// The mobile phone number.
	//
	// example:
	//
	// 1390000****
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// Indicates whether the mobile phone number supports card messages.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Support *bool `json:"support,omitempty" xml:"support,omitempty"`
}

func (s CheckMobilesCardSupportResponseBodyDataQueryResult) String() string {
	return dara.Prettify(s)
}

func (s CheckMobilesCardSupportResponseBodyDataQueryResult) GoString() string {
	return s.String()
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) GetMobile() *string {
	return s.Mobile
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) GetSupport() *bool {
	return s.Support
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) SetMobile(v string) *CheckMobilesCardSupportResponseBodyDataQueryResult {
	s.Mobile = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) SetSupport(v bool) *CheckMobilesCardSupportResponseBodyDataQueryResult {
	s.Support = &v
	return s
}

func (s *CheckMobilesCardSupportResponseBodyDataQueryResult) Validate() error {
	return dara.Validate(s)
}
