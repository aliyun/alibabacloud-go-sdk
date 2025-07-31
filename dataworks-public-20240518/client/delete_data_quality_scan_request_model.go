// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityScanRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DeleteDataQualityScanRequest
	GetDescription() *string
	SetId(v int64) *DeleteDataQualityScanRequest
	GetId() *int64
}

type DeleteDataQualityScanRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 123123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDataQualityScanRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityScanRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityScanRequest) GetDescription() *string {
	return s.Description
}

func (s *DeleteDataQualityScanRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDataQualityScanRequest) SetDescription(v string) *DeleteDataQualityScanRequest {
	s.Description = &v
	return s
}

func (s *DeleteDataQualityScanRequest) SetId(v int64) *DeleteDataQualityScanRequest {
	s.Id = &v
	return s
}

func (s *DeleteDataQualityScanRequest) Validate() error {
	return dara.Validate(s)
}
