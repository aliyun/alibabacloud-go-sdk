// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIngressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIngressId(v int64) *DeleteIngressRequest
	GetIngressId() *int64
}

type DeleteIngressRequest struct {
	// The ID of the routing rule that you want to delete. You can call the [ListIngresses](https://help.aliyun.com/document_detail/153934.html) operation to obtain the ID of a routing rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// 87
	IngressId *int64 `json:"IngressId,omitempty" xml:"IngressId,omitempty"`
}

func (s DeleteIngressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIngressRequest) GoString() string {
	return s.String()
}

func (s *DeleteIngressRequest) GetIngressId() *int64 {
	return s.IngressId
}

func (s *DeleteIngressRequest) SetIngressId(v int64) *DeleteIngressRequest {
	s.IngressId = &v
	return s
}

func (s *DeleteIngressRequest) Validate() error {
	return dara.Validate(s)
}
