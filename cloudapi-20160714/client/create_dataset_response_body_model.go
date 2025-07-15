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
	// The ID of the dataset.
	//
	// example:
	//
	// a25a6589b2584ff490e891cc********
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4173F95B-360C-460C-9F6C-4A96********
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
