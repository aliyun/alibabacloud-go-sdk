// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenVClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *OpenVClusterResponseBody
	GetData() *string
	SetRequestId(v string) *OpenVClusterResponseBody
	GetRequestId() *string
}

type OpenVClusterResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// rre59xelcx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 42E58E4D-ACAD-4400-8FAF-F762340AE5B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OpenVClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenVClusterResponseBody) GoString() string {
	return s.String()
}

func (s *OpenVClusterResponseBody) GetData() *string {
	return s.Data
}

func (s *OpenVClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenVClusterResponseBody) SetData(v string) *OpenVClusterResponseBody {
	s.Data = &v
	return s
}

func (s *OpenVClusterResponseBody) SetRequestId(v string) *OpenVClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenVClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
