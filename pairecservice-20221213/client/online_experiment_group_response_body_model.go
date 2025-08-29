// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnlineExperimentGroupResponseBody
	GetRequestId() *string
}

type OnlineExperimentGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 47F761ED-BE4E-51A6-B678-78E1490DF313
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineExperimentGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentGroupResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineExperimentGroupResponseBody) SetRequestId(v string) *OnlineExperimentGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineExperimentGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
