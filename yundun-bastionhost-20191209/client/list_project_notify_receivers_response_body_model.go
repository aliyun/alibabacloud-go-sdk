// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectNotifyReceiversResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*ListProjectNotifyReceiversResponseBodyItems) *ListProjectNotifyReceiversResponseBody
	GetItems() []*ListProjectNotifyReceiversResponseBodyItems
	SetMaxResults(v int32) *ListProjectNotifyReceiversResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListProjectNotifyReceiversResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListProjectNotifyReceiversResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListProjectNotifyReceiversResponseBody
	GetTotalCount() *int32
}

type ListProjectNotifyReceiversResponseBody struct {
	// The list of recipients.
	Items []*ListProjectNotifyReceiversResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The value of the MaxResults request parameter passed in the request if a valid parameter value was specified. Otherwise, the default value is returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next request. Use this value for the NextToken request parameter in the next paged query. If this field is empty in the response, no more data is available. This parameter is used for paging.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 81500666-d7f5-4143-8329-0223cc738105
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of recipients returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectNotifyReceiversResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectNotifyReceiversResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectNotifyReceiversResponseBody) GetItems() []*ListProjectNotifyReceiversResponseBodyItems {
	return s.Items
}

func (s *ListProjectNotifyReceiversResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListProjectNotifyReceiversResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListProjectNotifyReceiversResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectNotifyReceiversResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectNotifyReceiversResponseBody) SetItems(v []*ListProjectNotifyReceiversResponseBodyItems) *ListProjectNotifyReceiversResponseBody {
	s.Items = v
	return s
}

func (s *ListProjectNotifyReceiversResponseBody) SetMaxResults(v int32) *ListProjectNotifyReceiversResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBody) SetNextToken(v string) *ListProjectNotifyReceiversResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBody) SetRequestId(v string) *ListProjectNotifyReceiversResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBody) SetTotalCount(v int32) *ListProjectNotifyReceiversResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBody) Validate() error {
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

type ListProjectNotifyReceiversResponseBodyItems struct {
	// The email address of the recipient.
	//
	// example:
	//
	// 6291****@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the recipient.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The phone number of the recipient.
	//
	// example:
	//
	// 131****6412
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The ID of the notification recipient.
	//
	// example:
	//
	// 1
	ReceiverId *string `json:"ReceiverId,omitempty" xml:"ReceiverId,omitempty"`
}

func (s ListProjectNotifyReceiversResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListProjectNotifyReceiversResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListProjectNotifyReceiversResponseBodyItems) GetEmail() *string {
	return s.Email
}

func (s *ListProjectNotifyReceiversResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListProjectNotifyReceiversResponseBodyItems) GetPhone() *string {
	return s.Phone
}

func (s *ListProjectNotifyReceiversResponseBodyItems) GetReceiverId() *string {
	return s.ReceiverId
}

func (s *ListProjectNotifyReceiversResponseBodyItems) SetEmail(v string) *ListProjectNotifyReceiversResponseBodyItems {
	s.Email = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBodyItems) SetName(v string) *ListProjectNotifyReceiversResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBodyItems) SetPhone(v string) *ListProjectNotifyReceiversResponseBodyItems {
	s.Phone = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBodyItems) SetReceiverId(v string) *ListProjectNotifyReceiversResponseBodyItems {
	s.ReceiverId = &v
	return s
}

func (s *ListProjectNotifyReceiversResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
