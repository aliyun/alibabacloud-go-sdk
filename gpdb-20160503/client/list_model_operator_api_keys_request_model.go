// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListModelOperatorApiKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelOperatorApiKeysRequest
	GetPageSize() *int32
}

type ListModelOperatorApiKeysRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListModelOperatorApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorApiKeysRequest) GoString() string {
	return s.String()
}

func (s *ListModelOperatorApiKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelOperatorApiKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelOperatorApiKeysRequest) SetPageNumber(v int32) *ListModelOperatorApiKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelOperatorApiKeysRequest) SetPageSize(v int32) *ListModelOperatorApiKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelOperatorApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
