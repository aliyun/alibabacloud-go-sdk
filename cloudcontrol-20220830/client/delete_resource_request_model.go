// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteResourceRequest
	GetClientToken() *string
	SetFilter(v map[string]interface{}) *DeleteResourceRequest
	GetFilter() map[string]interface{}
	SetRegionId(v string) *DeleteResourceRequest
	GetRegionId() *string
}

type DeleteResourceRequest struct {
	// A client token to ensure idempotence. This parameter works only if the cloud product supports idempotence.
	//
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken *string                `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	Filter      map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
	// The region ID. This parameter is required if the cloud product is region-specific.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteResourceRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *DeleteResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteResourceRequest) SetClientToken(v string) *DeleteResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteResourceRequest) SetFilter(v map[string]interface{}) *DeleteResourceRequest {
	s.Filter = v
	return s
}

func (s *DeleteResourceRequest) SetRegionId(v string) *DeleteResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
