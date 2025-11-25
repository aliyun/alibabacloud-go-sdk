// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceSyncTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *CreateInstanceSyncTaskResponseBody
	GetModule() *string
	SetRequestId(v string) *CreateInstanceSyncTaskResponseBody
	GetRequestId() *string
}

type CreateInstanceSyncTaskResponseBody struct {
	// example:
	//
	// ips_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// D19D8F70-D64B-5A95-905A-6073BF4A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceSyncTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceSyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceSyncTaskResponseBody) GetModule() *string {
	return s.Module
}

func (s *CreateInstanceSyncTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceSyncTaskResponseBody) SetModule(v string) *CreateInstanceSyncTaskResponseBody {
	s.Module = &v
	return s
}

func (s *CreateInstanceSyncTaskResponseBody) SetRequestId(v string) *CreateInstanceSyncTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceSyncTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
