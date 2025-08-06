// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelMediaExportJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedJobIds(v []*string) *CancelMediaExportJobsResponseBody
	GetFailedJobIds() []*string
	SetNonExistJobIds(v []*string) *CancelMediaExportJobsResponseBody
	GetNonExistJobIds() []*string
	SetRequestId(v string) *CancelMediaExportJobsResponseBody
	GetRequestId() *string
}

type CancelMediaExportJobsResponseBody struct {
	FailedJobIds   []*string `json:"FailedJobIds,omitempty" xml:"FailedJobIds,omitempty" type:"Repeated"`
	NonExistJobIds []*string `json:"NonExistJobIds,omitempty" xml:"NonExistJobIds,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelMediaExportJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelMediaExportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *CancelMediaExportJobsResponseBody) GetFailedJobIds() []*string {
	return s.FailedJobIds
}

func (s *CancelMediaExportJobsResponseBody) GetNonExistJobIds() []*string {
	return s.NonExistJobIds
}

func (s *CancelMediaExportJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelMediaExportJobsResponseBody) SetFailedJobIds(v []*string) *CancelMediaExportJobsResponseBody {
	s.FailedJobIds = v
	return s
}

func (s *CancelMediaExportJobsResponseBody) SetNonExistJobIds(v []*string) *CancelMediaExportJobsResponseBody {
	s.NonExistJobIds = v
	return s
}

func (s *CancelMediaExportJobsResponseBody) SetRequestId(v string) *CancelMediaExportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelMediaExportJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
