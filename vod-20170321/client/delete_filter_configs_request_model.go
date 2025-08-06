// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilterConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUuId(v string) *DeleteFilterConfigsRequest
	GetUuId() *string
}

type DeleteFilterConfigsRequest struct {
	UuId *string `json:"UuId,omitempty" xml:"UuId,omitempty"`
}

func (s DeleteFilterConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilterConfigsRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilterConfigsRequest) GetUuId() *string {
	return s.UuId
}

func (s *DeleteFilterConfigsRequest) SetUuId(v string) *DeleteFilterConfigsRequest {
	s.UuId = &v
	return s
}

func (s *DeleteFilterConfigsRequest) Validate() error {
	return dara.Validate(s)
}
