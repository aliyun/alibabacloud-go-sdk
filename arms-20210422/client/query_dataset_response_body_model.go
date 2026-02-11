// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryDatasetResponseBody
	GetData() *string
	SetRequestId(v string) *QueryDatasetResponseBody
	GetRequestId() *string
}

type QueryDatasetResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDatasetResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDatasetResponseBody) SetData(v string) *QueryDatasetResponseBody {
	s.Data = &v
	return s
}

func (s *QueryDatasetResponseBody) SetRequestId(v string) *QueryDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
