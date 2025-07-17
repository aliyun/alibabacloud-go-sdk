// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteGrafanaResourceResponseBody
	GetCode() *int32
	SetData(v string) *DeleteGrafanaResourceResponseBody
	GetData() *string
	SetMessage(v string) *DeleteGrafanaResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGrafanaResourceResponseBody
	GetRequestId() *string
}

type DeleteGrafanaResourceResponseBody struct {
	// The status code. The status code 200 indicates a successful request, whereas others indicate a failed request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// delete success.
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 771DC66C-C5E0-59BC-A983-DD18FEE9EFFA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGrafanaResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteGrafanaResourceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteGrafanaResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGrafanaResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGrafanaResourceResponseBody) SetCode(v int32) *DeleteGrafanaResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetData(v string) *DeleteGrafanaResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetMessage(v string) *DeleteGrafanaResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetRequestId(v string) *DeleteGrafanaResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
