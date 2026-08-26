// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetId(v string) *CreateDatasetResponseBody
	GetDatasetId() *string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
}

type CreateDatasetResponseBody struct {
	// The dataset ID.
	//
	// example:
	//
	// d-rbvg5*****jhc9ks92
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B2C51F93-1C07-5477-9705-5FDB****F19F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetDatasetId() *string {
	return s.DatasetId
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) SetDatasetId(v string) *CreateDatasetResponseBody {
	s.DatasetId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
