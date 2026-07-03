// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *SetStorageResponseBody
	GetData() *bool
	SetRequestId(v string) *SetStorageResponseBody
	GetRequestId() *string
}

type SetStorageResponseBody struct {
	// Indicates whether the settings were saved. Valid values:
	//
	// - true: The settings were saved.
	//
	// - false: The settings failed to be saved.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-58D4-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *SetStorageResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetStorageResponseBody) SetData(v bool) *SetStorageResponseBody {
	s.Data = &v
	return s
}

func (s *SetStorageResponseBody) SetRequestId(v string) *SetStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
