// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateExperimentContentResponseBody
	GetRequestId() *string
	SetVersion(v int64) *UpdateExperimentContentResponseBody
	GetVersion() *int64
}

type UpdateExperimentContentResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateExperimentContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentContentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExperimentContentResponseBody) GetVersion() *int64 {
	return s.Version
}

func (s *UpdateExperimentContentResponseBody) SetRequestId(v string) *UpdateExperimentContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExperimentContentResponseBody) SetVersion(v int64) *UpdateExperimentContentResponseBody {
	s.Version = &v
	return s
}

func (s *UpdateExperimentContentResponseBody) Validate() error {
	return dara.Validate(s)
}
