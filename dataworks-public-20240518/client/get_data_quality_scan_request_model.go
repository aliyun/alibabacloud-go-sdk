// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityScanRequest
	GetId() *int64
}

type GetDataQualityScanRequest struct {
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityScanRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanRequest) SetId(v int64) *GetDataQualityScanRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRequest) Validate() error {
	return dara.Validate(s)
}
