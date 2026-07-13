// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupportedZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySupportedZonesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *QuerySupportedZonesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*QuerySupportedZonesResponseBodyItems) *QuerySupportedZonesResponseBody
	GetItems() []*QuerySupportedZonesResponseBodyItems
	SetMaxResults(v int32) *QuerySupportedZonesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *QuerySupportedZonesResponseBody
	GetMessage() *string
	SetNextToken(v string) *QuerySupportedZonesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *QuerySupportedZonesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySupportedZonesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *QuerySupportedZonesResponseBody
	GetTotalCount() *int32
}

type QuerySupportedZonesResponseBody struct {
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Items          []*QuerySupportedZonesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySupportedZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySupportedZonesResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySupportedZonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySupportedZonesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QuerySupportedZonesResponseBody) GetItems() []*QuerySupportedZonesResponseBodyItems {
	return s.Items
}

func (s *QuerySupportedZonesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QuerySupportedZonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySupportedZonesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *QuerySupportedZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySupportedZonesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySupportedZonesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySupportedZonesResponseBody) SetCode(v string) *QuerySupportedZonesResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetHttpStatusCode(v int32) *QuerySupportedZonesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetItems(v []*QuerySupportedZonesResponseBodyItems) *QuerySupportedZonesResponseBody {
	s.Items = v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetMaxResults(v int32) *QuerySupportedZonesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetMessage(v string) *QuerySupportedZonesResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetNextToken(v string) *QuerySupportedZonesResponseBody {
	s.NextToken = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetRequestId(v string) *QuerySupportedZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetSuccess(v bool) *QuerySupportedZonesResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) SetTotalCount(v int32) *QuerySupportedZonesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QuerySupportedZonesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QuerySupportedZonesResponseBodyItems struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s QuerySupportedZonesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySupportedZonesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *QuerySupportedZonesResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *QuerySupportedZonesResponseBodyItems) SetZoneId(v string) *QuerySupportedZonesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *QuerySupportedZonesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
