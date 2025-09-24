// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetJobConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetJobConfigId(v string) *CreateDatasetJobConfigResponseBody
	GetDatasetJobConfigId() *string
	SetRequestId(v string) *CreateDatasetJobConfigResponseBody
	GetRequestId() *string
}

type CreateDatasetJobConfigResponseBody struct {
	// The configuration ID.
	//
	// example:
	//
	// dscfg-xxxxxxxxxxxxxx
	DatasetJobConfigId *string `json:"DatasetJobConfigId,omitempty" xml:"DatasetJobConfigId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetJobConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetJobConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetJobConfigResponseBody) GetDatasetJobConfigId() *string {
	return s.DatasetJobConfigId
}

func (s *CreateDatasetJobConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetJobConfigResponseBody) SetDatasetJobConfigId(v string) *CreateDatasetJobConfigResponseBody {
	s.DatasetJobConfigId = &v
	return s
}

func (s *CreateDatasetJobConfigResponseBody) SetRequestId(v string) *CreateDatasetJobConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetJobConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
