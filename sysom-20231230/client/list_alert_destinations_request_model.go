// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAlertDestinationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetXDebugId(v string) *ListAlertDestinationsRequest
	GetXDebugId() *string
	SetCurrent(v int32) *ListAlertDestinationsRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListAlertDestinationsRequest
	GetMaxResults() *int32
	SetName(v string) *ListAlertDestinationsRequest
	GetName() *string
	SetNextToken(v string) *ListAlertDestinationsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListAlertDestinationsRequest
	GetPageSize() *int32
	SetXSysomInvokeSource(v string) *ListAlertDestinationsRequest
	GetXSysomInvokeSource() *string
}

type ListAlertDestinationsRequest struct {
	XDebugId *string `json:"X-Debug-Id,omitempty" xml:"X-Debug-Id,omitempty"`
	// The current page number (starting from 1).
	//
	// example:
	//
	// 1
	Current *int32 `json:"current,omitempty" xml:"current,omitempty"`
	// The maximum number of records to retrieve in a single request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// name1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token for the next request.
	//
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize           *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	XSysomInvokeSource *string `json:"x-sysom-invoke-source,omitempty" xml:"x-sysom-invoke-source,omitempty"`
}

func (s ListAlertDestinationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAlertDestinationsRequest) GoString() string {
	return s.String()
}

func (s *ListAlertDestinationsRequest) GetXDebugId() *string {
	return s.XDebugId
}

func (s *ListAlertDestinationsRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAlertDestinationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAlertDestinationsRequest) GetName() *string {
	return s.Name
}

func (s *ListAlertDestinationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAlertDestinationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAlertDestinationsRequest) GetXSysomInvokeSource() *string {
	return s.XSysomInvokeSource
}

func (s *ListAlertDestinationsRequest) SetXDebugId(v string) *ListAlertDestinationsRequest {
	s.XDebugId = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetCurrent(v int32) *ListAlertDestinationsRequest {
	s.Current = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetMaxResults(v int32) *ListAlertDestinationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetName(v string) *ListAlertDestinationsRequest {
	s.Name = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetNextToken(v string) *ListAlertDestinationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetPageSize(v int32) *ListAlertDestinationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAlertDestinationsRequest) SetXSysomInvokeSource(v string) *ListAlertDestinationsRequest {
	s.XSysomInvokeSource = &v
	return s
}

func (s *ListAlertDestinationsRequest) Validate() error {
	return dara.Validate(s)
}
