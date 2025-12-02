// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateBackupConfigResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateBackupConfigResponseBody
	GetRequestId() *string
}

type UpdateBackupConfigResponseBody struct {
	// Returned data.
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

func (s UpdateBackupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateBackupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupConfigResponseBody) SetData(v bool) *UpdateBackupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBackupConfigResponseBody) SetRequestId(v string) *UpdateBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
