// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPropertyId(v int64) *RemovePropertyRequest
	GetPropertyId() *int64
}

type RemovePropertyRequest struct {
	// The ID of the property. You can call the [ListProperty](https://help.aliyun.com/document_detail/410890.html) operation to query the property ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 390
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
}

func (s RemovePropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s RemovePropertyRequest) GoString() string {
	return s.String()
}

func (s *RemovePropertyRequest) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *RemovePropertyRequest) SetPropertyId(v int64) *RemovePropertyRequest {
	s.PropertyId = &v
	return s
}

func (s *RemovePropertyRequest) Validate() error {
	return dara.Validate(s)
}
