// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeColdStorageSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ResizeColdStorageSizeRequest
	GetClusterId() *string
	SetColdStorageSize(v int32) *ResizeColdStorageSizeRequest
	GetColdStorageSize() *int32
}

type ResizeColdStorageSizeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ld-bp169l540vc6c****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 900
	ColdStorageSize *int32 `json:"ColdStorageSize,omitempty" xml:"ColdStorageSize,omitempty"`
}

func (s ResizeColdStorageSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ResizeColdStorageSizeRequest) GoString() string {
	return s.String()
}

func (s *ResizeColdStorageSizeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ResizeColdStorageSizeRequest) GetColdStorageSize() *int32 {
	return s.ColdStorageSize
}

func (s *ResizeColdStorageSizeRequest) SetClusterId(v string) *ResizeColdStorageSizeRequest {
	s.ClusterId = &v
	return s
}

func (s *ResizeColdStorageSizeRequest) SetColdStorageSize(v int32) *ResizeColdStorageSizeRequest {
	s.ColdStorageSize = &v
	return s
}

func (s *ResizeColdStorageSizeRequest) Validate() error {
	return dara.Validate(s)
}
