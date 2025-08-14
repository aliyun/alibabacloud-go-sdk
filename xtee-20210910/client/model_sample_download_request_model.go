// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSampleDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegId(v string) *ModelSampleDownloadRequest
	GetRegId() *string
}

type ModelSampleDownloadRequest struct {
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
}

func (s ModelSampleDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelSampleDownloadRequest) GoString() string {
	return s.String()
}

func (s *ModelSampleDownloadRequest) GetRegId() *string {
	return s.RegId
}

func (s *ModelSampleDownloadRequest) SetRegId(v string) *ModelSampleDownloadRequest {
	s.RegId = &v
	return s
}

func (s *ModelSampleDownloadRequest) Validate() error {
	return dara.Validate(s)
}
