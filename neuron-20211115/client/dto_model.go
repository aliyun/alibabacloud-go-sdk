// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDTO interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DTO
	GetId() *int64
	SetRequestId(v string) *DTO
	GetRequestId() *string
}

type DTO struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DTO) String() string {
	return dara.Prettify(s)
}

func (s DTO) GoString() string {
	return s.String()
}

func (s *DTO) GetId() *int64 {
	return s.Id
}

func (s *DTO) GetRequestId() *string {
	return s.RequestId
}

func (s *DTO) SetId(v int64) *DTO {
	s.Id = &v
	return s
}

func (s *DTO) SetRequestId(v string) *DTO {
	s.RequestId = &v
	return s
}

func (s *DTO) Validate() error {
	return dara.Validate(s)
}
