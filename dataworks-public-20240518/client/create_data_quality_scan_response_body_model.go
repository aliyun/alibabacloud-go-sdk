// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityScanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityScanResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityScanResponseBody
	GetRequestId() *string
}

type CreateDataQualityScanResponseBody struct {
	// example:
	//
	// 676303114031776
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityScanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityScanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityScanResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityScanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityScanResponseBody) SetId(v int64) *CreateDataQualityScanResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityScanResponseBody) SetRequestId(v string) *CreateDataQualityScanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityScanResponseBody) Validate() error {
	return dara.Validate(s)
}
