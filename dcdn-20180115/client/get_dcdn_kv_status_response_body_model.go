// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComplete(v bool) *GetDcdnKvStatusResponseBody
	GetComplete() *bool
	SetExpire(v string) *GetDcdnKvStatusResponseBody
	GetExpire() *string
	SetRequestId(v string) *GetDcdnKvStatusResponseBody
	GetRequestId() *string
}

type GetDcdnKvStatusResponseBody struct {
	// Specifies whether the configured key has taken effect on all points of presence (POPs).
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// The timeout period of the configured key. The value is an absolute timestamp, such as 2023-09-11T15:39:44+08:00. This parameter is not returned if the key is permanently stored.
	//
	// example:
	//
	// 2023-09-11T15:39:44+08:00
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDcdnKvStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetDcdnKvStatusResponseBody) GetComplete() *bool {
	return s.Complete
}

func (s *GetDcdnKvStatusResponseBody) GetExpire() *string {
	return s.Expire
}

func (s *GetDcdnKvStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDcdnKvStatusResponseBody) SetComplete(v bool) *GetDcdnKvStatusResponseBody {
	s.Complete = &v
	return s
}

func (s *GetDcdnKvStatusResponseBody) SetExpire(v string) *GetDcdnKvStatusResponseBody {
	s.Expire = &v
	return s
}

func (s *GetDcdnKvStatusResponseBody) SetRequestId(v string) *GetDcdnKvStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDcdnKvStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
