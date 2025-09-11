// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateDatasetResponseBody
	GetResult() *string
	SetSuccess(v bool) *CreateDatasetResponseBody
	GetSuccess() *bool
}

type CreateDatasetResponseBody struct {
	// example:
	//
	// F68B***********A3DF743
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) SetResult(v string) *CreateDatasetResponseBody {
	s.Result = &v
	return s
}

func (s *CreateDatasetResponseBody) SetSuccess(v bool) *CreateDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
