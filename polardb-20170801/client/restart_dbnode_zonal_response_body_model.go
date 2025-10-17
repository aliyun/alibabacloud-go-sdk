// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBNodeZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDBNodeZonalResponseBody
	GetRequestId() *string
}

type RestartDBNodeZonalResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBNodeZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBNodeZonalResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBNodeZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBNodeZonalResponseBody) SetRequestId(v string) *RestartDBNodeZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBNodeZonalResponseBody) Validate() error {
	return dara.Validate(s)
}
