// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnlineExperimentResponseBody
	GetRequestId() *string
}

type OnlineExperimentResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6CF1E160-3F36-5E73-A170-C75504F05BBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineExperimentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineExperimentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineExperimentResponseBody) SetRequestId(v string) *OnlineExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineExperimentResponseBody) Validate() error {
	return dara.Validate(s)
}
