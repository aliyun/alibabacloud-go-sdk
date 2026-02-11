// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGrafanaResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteGrafanaResourceResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteGrafanaResourceResponseBody
	GetRequestId() *string
}

type DeleteGrafanaResourceResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGrafanaResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGrafanaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGrafanaResourceResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteGrafanaResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGrafanaResourceResponseBody) SetData(v string) *DeleteGrafanaResourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) SetRequestId(v string) *DeleteGrafanaResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGrafanaResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
