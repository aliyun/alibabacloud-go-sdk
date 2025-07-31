// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityScanRunRequest
	GetId() *int64
}

type GetDataQualityScanRunRequest struct {
	// example:
	//
	// 1006059507
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityScanRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanRunRequest) SetId(v int64) *GetDataQualityScanRunRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRunRequest) Validate() error {
	return dara.Validate(s)
}
