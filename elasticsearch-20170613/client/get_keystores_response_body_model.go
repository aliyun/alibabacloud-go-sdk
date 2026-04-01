// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeystoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetKeystoresResponseBody
	GetRequestId() *string
	SetResult(v []*string) *GetKeystoresResponseBody
	GetResult() []*string
}

type GetKeystoresResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC****
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetKeystoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKeystoresResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeystoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKeystoresResponseBody) GetResult() []*string {
	return s.Result
}

func (s *GetKeystoresResponseBody) SetRequestId(v string) *GetKeystoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeystoresResponseBody) SetResult(v []*string) *GetKeystoresResponseBody {
	s.Result = v
	return s
}

func (s *GetKeystoresResponseBody) Validate() error {
	return dara.Validate(s)
}
