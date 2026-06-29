// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRevokeSeatsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*BatchRevokeSeatsRequestItems) *BatchRevokeSeatsRequest
	GetItems() []*BatchRevokeSeatsRequestItems
	SetLocale(v string) *BatchRevokeSeatsRequest
	GetLocale() *string
}

type BatchRevokeSeatsRequest struct {
	// The list of revocation items. This parameter is required.
	Items []*BatchRevokeSeatsRequestItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The language. Valid values: zh-CN and en-US.
	//
	// example:
	//
	// zh-CN
	Locale *string `json:"Locale,omitempty" xml:"Locale,omitempty"`
}

func (s BatchRevokeSeatsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokeSeatsRequest) GoString() string {
	return s.String()
}

func (s *BatchRevokeSeatsRequest) GetItems() []*BatchRevokeSeatsRequestItems {
	return s.Items
}

func (s *BatchRevokeSeatsRequest) GetLocale() *string {
	return s.Locale
}

func (s *BatchRevokeSeatsRequest) SetItems(v []*BatchRevokeSeatsRequestItems) *BatchRevokeSeatsRequest {
	s.Items = v
	return s
}

func (s *BatchRevokeSeatsRequest) SetLocale(v string) *BatchRevokeSeatsRequest {
	s.Locale = &v
	return s
}

func (s *BatchRevokeSeatsRequest) Validate() error {
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

type BatchRevokeSeatsRequestItems struct {
	// The ID of the currently associated member.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s BatchRevokeSeatsRequestItems) String() string {
	return dara.Prettify(s)
}

func (s BatchRevokeSeatsRequestItems) GoString() string {
	return s.String()
}

func (s *BatchRevokeSeatsRequestItems) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchRevokeSeatsRequestItems) SetAccountId(v string) *BatchRevokeSeatsRequestItems {
	s.AccountId = &v
	return s
}

func (s *BatchRevokeSeatsRequestItems) Validate() error {
	return dara.Validate(s)
}
