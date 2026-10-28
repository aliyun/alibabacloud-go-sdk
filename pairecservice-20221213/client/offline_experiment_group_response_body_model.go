// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineExperimentGroupResponseBody
	GetRequestId() *string
}

type OfflineExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 518C64F6-DFF7-11ED-85B0-00163E14B3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineExperimentGroupResponseBody) SetRequestId(v string) *OfflineExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
