// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaExportJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedJobIds(v []*string) *DeleteMediaExportJobsResponseBody
	GetFailedJobIds() []*string
	SetNonExistJobIds(v []*string) *DeleteMediaExportJobsResponseBody
	GetNonExistJobIds() []*string
	SetRequestId(v string) *DeleteMediaExportJobsResponseBody
	GetRequestId() *string
}

type DeleteMediaExportJobsResponseBody struct {
	FailedJobIds   []*string `json:"FailedJobIds,omitempty" xml:"FailedJobIds,omitempty" type:"Repeated"`
	NonExistJobIds []*string `json:"NonExistJobIds,omitempty" xml:"NonExistJobIds,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaExportJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaExportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaExportJobsResponseBody) GetFailedJobIds() []*string {
	return s.FailedJobIds
}

func (s *DeleteMediaExportJobsResponseBody) GetNonExistJobIds() []*string {
	return s.NonExistJobIds
}

func (s *DeleteMediaExportJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaExportJobsResponseBody) SetFailedJobIds(v []*string) *DeleteMediaExportJobsResponseBody {
	s.FailedJobIds = v
	return s
}

func (s *DeleteMediaExportJobsResponseBody) SetNonExistJobIds(v []*string) *DeleteMediaExportJobsResponseBody {
	s.NonExistJobIds = v
	return s
}

func (s *DeleteMediaExportJobsResponseBody) SetRequestId(v string) *DeleteMediaExportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaExportJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
