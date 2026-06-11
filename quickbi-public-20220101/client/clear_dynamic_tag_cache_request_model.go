// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearDynamicTagCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ClearDynamicTagCacheRequest
	GetConfigId() *string
}

type ClearDynamicTagCacheRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cfg************405791744
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
}

func (s ClearDynamicTagCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearDynamicTagCacheRequest) GoString() string {
	return s.String()
}

func (s *ClearDynamicTagCacheRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ClearDynamicTagCacheRequest) SetConfigId(v string) *ClearDynamicTagCacheRequest {
	s.ConfigId = &v
	return s
}

func (s *ClearDynamicTagCacheRequest) Validate() error {
	return dara.Validate(s)
}
