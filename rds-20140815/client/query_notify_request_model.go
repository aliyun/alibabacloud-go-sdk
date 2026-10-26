// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryNotifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *QueryNotifyRequest
	GetFrom() *string
	SetPageNumber(v int32) *QueryNotifyRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryNotifyRequest
	GetPageSize() *int32
	SetTo(v string) *QueryNotifyRequest
	GetTo() *string
	SetWithConfirmed(v bool) *QueryNotifyRequest
	GetWithConfirmed() *bool
}

type QueryNotifyRequest struct {
	// The start time of the query. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-02T08:38:37Z
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The page number. Valid values: integers greater than **0*	- and less than or equal to the maximum value of the Integer type.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **30**.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The end time of the query. The end time must be later than the start time. Specify the time in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-05-09T08:38:37Z
	To *string `json:"To,omitempty" xml:"To,omitempty"`
	// Specifies whether the query results include confirmed notifications. Valid values:
	//
	// - **true**: Includes confirmed notifications.
	//
	// - **false**: Does not include confirmed notifications.
	//
	// > Confirmed notifications are those that have been marked as confirmed by calling the ConfirmNotify operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	WithConfirmed *bool `json:"WithConfirmed,omitempty" xml:"WithConfirmed,omitempty"`
}

func (s QueryNotifyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryNotifyRequest) GoString() string {
	return s.String()
}

func (s *QueryNotifyRequest) GetFrom() *string {
	return s.From
}

func (s *QueryNotifyRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryNotifyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryNotifyRequest) GetTo() *string {
	return s.To
}

func (s *QueryNotifyRequest) GetWithConfirmed() *bool {
	return s.WithConfirmed
}

func (s *QueryNotifyRequest) SetFrom(v string) *QueryNotifyRequest {
	s.From = &v
	return s
}

func (s *QueryNotifyRequest) SetPageNumber(v int32) *QueryNotifyRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryNotifyRequest) SetPageSize(v int32) *QueryNotifyRequest {
	s.PageSize = &v
	return s
}

func (s *QueryNotifyRequest) SetTo(v string) *QueryNotifyRequest {
	s.To = &v
	return s
}

func (s *QueryNotifyRequest) SetWithConfirmed(v bool) *QueryNotifyRequest {
	s.WithConfirmed = &v
	return s
}

func (s *QueryNotifyRequest) Validate() error {
	return dara.Validate(s)
}
