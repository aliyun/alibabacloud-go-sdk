// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineExperimentResponseBody
	GetRequestId() *string
}

type OfflineExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 872951C9-7755-5FA1-AACD-7F9375A6D27A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineExperimentResponseBody) SetRequestId(v string) *OfflineExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
