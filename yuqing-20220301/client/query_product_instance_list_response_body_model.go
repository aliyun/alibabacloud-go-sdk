// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProductInstanceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceList(v []*ProductInstance) *QueryProductInstanceListResponseBody
	GetInstanceList() []*ProductInstance
	SetRequestId(v string) *QueryProductInstanceListResponseBody
	GetRequestId() *string
}

type QueryProductInstanceListResponseBody struct {
	InstanceList []*ProductInstance `json:"instanceList,omitempty" xml:"instanceList,omitempty" type:"Repeated"`
	RequestId    *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryProductInstanceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProductInstanceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProductInstanceListResponseBody) GetInstanceList() []*ProductInstance {
	return s.InstanceList
}

func (s *QueryProductInstanceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProductInstanceListResponseBody) SetInstanceList(v []*ProductInstance) *QueryProductInstanceListResponseBody {
	s.InstanceList = v
	return s
}

func (s *QueryProductInstanceListResponseBody) SetRequestId(v string) *QueryProductInstanceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProductInstanceListResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
