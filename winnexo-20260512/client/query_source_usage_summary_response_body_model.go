// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySourceUsageSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySourceUsageSummaryResponseBody
	GetCode() *string
	SetMessage(v string) *QuerySourceUsageSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySourceUsageSummaryResponseBody
	GetRequestId() *string
	SetRetrievalCount(v int64) *QuerySourceUsageSummaryResponseBody
	GetRetrievalCount() *int64
	SetRetrievedSourceCount(v int64) *QuerySourceUsageSummaryResponseBody
	GetRetrievedSourceCount() *int64
	SetSourceCount(v int64) *QuerySourceUsageSummaryResponseBody
	GetSourceCount() *int64
	SetUsageCount(v int64) *QuerySourceUsageSummaryResponseBody
	GetUsageCount() *int64
	SetUsedSourceCount(v int64) *QuerySourceUsageSummaryResponseBody
	GetUsedSourceCount() *int64
}

type QuerySourceUsageSummaryResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of retrieval recalls.
	//
	// example:
	//
	// 10
	RetrievalCount *int64 `json:"retrievalCount,omitempty" xml:"retrievalCount,omitempty"`
	// The deduplicated count of sources that have retrieval events.
	//
	// example:
	//
	// 1
	RetrievedSourceCount *int64 `json:"retrievedSourceCount,omitempty" xml:"retrievedSourceCount,omitempty"`
	// The number of rows in the source data.
	//
	// example:
	//
	// 10
	SourceCount *int64 `json:"sourceCount,omitempty" xml:"sourceCount,omitempty"`
	// The number of times the source has been used in agents.
	//
	// example:
	//
	// 8
	UsageCount *int64 `json:"usageCount,omitempty" xml:"usageCount,omitempty"`
	// The deduplicated count of sources that have usage events.
	//
	// example:
	//
	// 2
	UsedSourceCount *int64 `json:"usedSourceCount,omitempty" xml:"usedSourceCount,omitempty"`
}

func (s QuerySourceUsageSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySourceUsageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySourceUsageSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySourceUsageSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySourceUsageSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySourceUsageSummaryResponseBody) GetRetrievalCount() *int64 {
	return s.RetrievalCount
}

func (s *QuerySourceUsageSummaryResponseBody) GetRetrievedSourceCount() *int64 {
	return s.RetrievedSourceCount
}

func (s *QuerySourceUsageSummaryResponseBody) GetSourceCount() *int64 {
	return s.SourceCount
}

func (s *QuerySourceUsageSummaryResponseBody) GetUsageCount() *int64 {
	return s.UsageCount
}

func (s *QuerySourceUsageSummaryResponseBody) GetUsedSourceCount() *int64 {
	return s.UsedSourceCount
}

func (s *QuerySourceUsageSummaryResponseBody) SetCode(v string) *QuerySourceUsageSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetMessage(v string) *QuerySourceUsageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetRequestId(v string) *QuerySourceUsageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetRetrievalCount(v int64) *QuerySourceUsageSummaryResponseBody {
	s.RetrievalCount = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetRetrievedSourceCount(v int64) *QuerySourceUsageSummaryResponseBody {
	s.RetrievedSourceCount = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetSourceCount(v int64) *QuerySourceUsageSummaryResponseBody {
	s.SourceCount = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetUsageCount(v int64) *QuerySourceUsageSummaryResponseBody {
	s.UsageCount = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) SetUsedSourceCount(v int64) *QuerySourceUsageSummaryResponseBody {
	s.UsedSourceCount = &v
	return s
}

func (s *QuerySourceUsageSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
