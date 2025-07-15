// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetItemId(v string) *CreateDatasetItemResponseBody
	GetDatasetItemId() *string
	SetRequestId(v string) *CreateDatasetItemResponseBody
	GetRequestId() *string
}

type CreateDatasetItemResponseBody struct {
	// example:
	//
	// 5045****
	DatasetItemId *string `json:"DatasetItemId,omitempty" xml:"DatasetItemId,omitempty"`
	// example:
	//
	// 8A5E2053-4D9F-5280-B7A9-D357********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetItemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetItemResponseBody) GetDatasetItemId() *string {
	return s.DatasetItemId
}

func (s *CreateDatasetItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetItemResponseBody) SetDatasetItemId(v string) *CreateDatasetItemResponseBody {
	s.DatasetItemId = &v
	return s
}

func (s *CreateDatasetItemResponseBody) SetRequestId(v string) *CreateDatasetItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetItemResponseBody) Validate() error {
	return dara.Validate(s)
}
