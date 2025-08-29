// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineLaboratoryResponseBody
	GetRequestId() *string
}

type OfflineLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineLaboratoryResponseBody) SetRequestId(v string) *OfflineLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
