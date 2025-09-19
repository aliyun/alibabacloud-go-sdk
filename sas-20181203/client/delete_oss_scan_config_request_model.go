// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOssScanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteOssScanConfigRequest
	GetId() *int64
}

type DeleteOssScanConfigRequest struct {
	// The configuration ID.
	//
	// >  You can call the [ListOssScanConfig](~~ListOssScanConfig~~) operation to query configuration IDs.
	//
	// example:
	//
	// 1589
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteOssScanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOssScanConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteOssScanConfigRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteOssScanConfigRequest) SetId(v int64) *DeleteOssScanConfigRequest {
	s.Id = &v
	return s
}

func (s *DeleteOssScanConfigRequest) Validate() error {
	return dara.Validate(s)
}
