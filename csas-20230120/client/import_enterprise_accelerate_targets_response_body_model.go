// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportEnterpriseAccelerateTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportEnterpriseAccelerateTargetsResponseBody
	GetData() *string
	SetRequestId(v string) *ImportEnterpriseAccelerateTargetsResponseBody
	GetRequestId() *string
}

type ImportEnterpriseAccelerateTargetsResponseBody struct {
	// example:
	//
	// 1648723859058501
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 305508BD-8A31-5E15-86CE-52D57967C45E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportEnterpriseAccelerateTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportEnterpriseAccelerateTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportEnterpriseAccelerateTargetsResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportEnterpriseAccelerateTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportEnterpriseAccelerateTargetsResponseBody) SetData(v string) *ImportEnterpriseAccelerateTargetsResponseBody {
	s.Data = &v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsResponseBody) SetRequestId(v string) *ImportEnterpriseAccelerateTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportEnterpriseAccelerateTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}
