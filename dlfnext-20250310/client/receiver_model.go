// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiver interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Receiver
	GetComment() *string
	SetCreatedAt(v int64) *Receiver
	GetCreatedAt() *int64
	SetCreatedBy(v string) *Receiver
	GetCreatedBy() *string
	SetReceiverName(v string) *Receiver
	GetReceiverName() *string
	SetReceiverTenantId(v int64) *Receiver
	GetReceiverTenantId() *int64
	SetUpdatedAt(v int64) *Receiver
	GetUpdatedAt() *int64
	SetUpdatedBy(v string) *Receiver
	GetUpdatedBy() *string
}

type Receiver struct {
	// example:
	//
	// comment
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// example:
	//
	// 1744970111419
	CreatedAt *int64 `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	CreatedBy *string `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// example:
	//
	// receiver_name
	ReceiverName *string `json:"receiverName,omitempty" xml:"receiverName,omitempty"`
	// example:
	//
	// 1111
	ReceiverTenantId *int64 `json:"receiverTenantId,omitempty" xml:"receiverTenantId,omitempty"`
	// example:
	//
	// 1744970111419
	UpdatedAt *int64 `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// acs:ram::[accountId]:root
	UpdatedBy *string `json:"updatedBy,omitempty" xml:"updatedBy,omitempty"`
}

func (s Receiver) String() string {
	return dara.Prettify(s)
}

func (s Receiver) GoString() string {
	return s.String()
}

func (s *Receiver) GetComment() *string {
	return s.Comment
}

func (s *Receiver) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *Receiver) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *Receiver) GetReceiverName() *string {
	return s.ReceiverName
}

func (s *Receiver) GetReceiverTenantId() *int64 {
	return s.ReceiverTenantId
}

func (s *Receiver) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *Receiver) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *Receiver) SetComment(v string) *Receiver {
	s.Comment = &v
	return s
}

func (s *Receiver) SetCreatedAt(v int64) *Receiver {
	s.CreatedAt = &v
	return s
}

func (s *Receiver) SetCreatedBy(v string) *Receiver {
	s.CreatedBy = &v
	return s
}

func (s *Receiver) SetReceiverName(v string) *Receiver {
	s.ReceiverName = &v
	return s
}

func (s *Receiver) SetReceiverTenantId(v int64) *Receiver {
	s.ReceiverTenantId = &v
	return s
}

func (s *Receiver) SetUpdatedAt(v int64) *Receiver {
	s.UpdatedAt = &v
	return s
}

func (s *Receiver) SetUpdatedBy(v string) *Receiver {
	s.UpdatedBy = &v
	return s
}

func (s *Receiver) Validate() error {
	return dara.Validate(s)
}
