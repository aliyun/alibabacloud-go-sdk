// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *GetBackupStatusResponseBody
	GetData() *bool
	SetRequestId(v string) *GetBackupStatusResponseBody
	GetRequestId() *string
}

type GetBackupStatusResponseBody struct {
	// Audit result data.
	//
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetBackupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetBackupStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *GetBackupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBackupStatusResponseBody) SetData(v bool) *GetBackupStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetBackupStatusResponseBody) SetRequestId(v string) *GetBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBackupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
