// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListenerHealthStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*GetListenerHealthStatusRequestFilter) *GetListenerHealthStatusRequest
	GetFilter() []*GetListenerHealthStatusRequestFilter
	SetListenerId(v string) *GetListenerHealthStatusRequest
	GetListenerId() *string
	SetMaxResults(v int32) *GetListenerHealthStatusRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetListenerHealthStatusRequest
	GetNextToken() *string
	SetSkip(v int32) *GetListenerHealthStatusRequest
	GetSkip() *int32
}

type GetListenerHealthStatusRequest struct {
	// The filter conditions. You can specify at most 20 filter conditions.
	Filter []*GetListenerHealthStatusRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The listener ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsn-7sixpvm5fc3v0b****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The number of entries per page. Valid values: 1 to 1000. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// 4f1d7cc9f51e18904e8a063603a6b0c3d03bc69f78734254e0b5e8707e68****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the call.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
}

func (s GetListenerHealthStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusRequest) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequest) GetFilter() []*GetListenerHealthStatusRequestFilter {
	return s.Filter
}

func (s *GetListenerHealthStatusRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *GetListenerHealthStatusRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetListenerHealthStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetListenerHealthStatusRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *GetListenerHealthStatusRequest) SetFilter(v []*GetListenerHealthStatusRequestFilter) *GetListenerHealthStatusRequest {
	s.Filter = v
	return s
}

func (s *GetListenerHealthStatusRequest) SetListenerId(v string) *GetListenerHealthStatusRequest {
	s.ListenerId = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetMaxResults(v int32) *GetListenerHealthStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetNextToken(v string) *GetListenerHealthStatusRequest {
	s.NextToken = &v
	return s
}

func (s *GetListenerHealthStatusRequest) SetSkip(v int32) *GetListenerHealthStatusRequest {
	s.Skip = &v
	return s
}

func (s *GetListenerHealthStatusRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetListenerHealthStatusRequestFilter struct {
	// The filter condition name. You can filter by one or more filter condition names. The URL must meet the following requirements:
	//
	// 	- **Status**: the health status.
	//
	// 	- **ReasonCode**: the cause of an unhealthy server.
	//
	// 	- **ServerId**: the ID of the backend server.
	//
	// 	- **ServerIp**: the IP address of the backend server.
	//
	// example:
	//
	// Status
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter condition values. You can specify at most 20 condition values.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s GetListenerHealthStatusRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s GetListenerHealthStatusRequestFilter) GoString() string {
	return s.String()
}

func (s *GetListenerHealthStatusRequestFilter) GetName() *string {
	return s.Name
}

func (s *GetListenerHealthStatusRequestFilter) GetValues() []*string {
	return s.Values
}

func (s *GetListenerHealthStatusRequestFilter) SetName(v string) *GetListenerHealthStatusRequestFilter {
	s.Name = &v
	return s
}

func (s *GetListenerHealthStatusRequestFilter) SetValues(v []*string) *GetListenerHealthStatusRequestFilter {
	s.Values = v
	return s
}

func (s *GetListenerHealthStatusRequestFilter) Validate() error {
	return dara.Validate(s)
}
