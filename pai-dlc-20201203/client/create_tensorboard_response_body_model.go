// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTensorboardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceId(v string) *CreateTensorboardResponseBody
	GetDataSourceId() *string
	SetJobId(v string) *CreateTensorboardResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateTensorboardResponseBody
	GetRequestId() *string
	SetTensorboardId(v string) *CreateTensorboardResponseBody
	GetTensorboardId() *string
}

type CreateTensorboardResponseBody struct {
	// The dataset ID.
	//
	// example:
	//
	// ds-20210126170216-xxxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dlc-20210126170216-xxxxxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// TensorBoard ID
	//
	// example:
	//
	// tbxxxxxxxx
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s CreateTensorboardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponseBody) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateTensorboardResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateTensorboardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTensorboardResponseBody) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *CreateTensorboardResponseBody) SetDataSourceId(v string) *CreateTensorboardResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetJobId(v string) *CreateTensorboardResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetRequestId(v string) *CreateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetTensorboardId(v string) *CreateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *CreateTensorboardResponseBody) Validate() error {
	return dara.Validate(s)
}
