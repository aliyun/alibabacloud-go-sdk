// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEdgeContainerAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateEdgeContainerAppResponseBody
	GetAppId() *string
	SetRequestId(v string) *CreateEdgeContainerAppResponseBody
	GetRequestId() *string
}

type CreateEdgeContainerAppResponseBody struct {
	// The ID of the application that is created.
	//
	// example:
	//
	// app-880688675783794688
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEdgeContainerAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEdgeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEdgeContainerAppResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *CreateEdgeContainerAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEdgeContainerAppResponseBody) SetAppId(v string) *CreateEdgeContainerAppResponseBody {
	s.AppId = &v
	return s
}

func (s *CreateEdgeContainerAppResponseBody) SetRequestId(v string) *CreateEdgeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEdgeContainerAppResponseBody) Validate() error {
	return dara.Validate(s)
}
