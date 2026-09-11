// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoreAvailableVscsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListKVCacheStoreAvailableVscsResponseBody
	GetRequestId() *string
	SetVscs(v []*ListKVCacheStoreAvailableVscsResponseBodyVscs) *ListKVCacheStoreAvailableVscsResponseBody
	GetVscs() []*ListKVCacheStoreAvailableVscsResponseBodyVscs
}

type ListKVCacheStoreAvailableVscsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// request-id-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of available VSC resources.
	Vscs []*ListKVCacheStoreAvailableVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s ListKVCacheStoreAvailableVscsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableVscsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableVscsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKVCacheStoreAvailableVscsResponseBody) GetVscs() []*ListKVCacheStoreAvailableVscsResponseBodyVscs {
	return s.Vscs
}

func (s *ListKVCacheStoreAvailableVscsResponseBody) SetRequestId(v string) *ListKVCacheStoreAvailableVscsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponseBody) SetVscs(v []*ListKVCacheStoreAvailableVscsResponseBodyVscs) *ListKVCacheStoreAvailableVscsResponseBody {
	s.Vscs = v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponseBody) Validate() error {
	if s.Vscs != nil {
		for _, item := range s.Vscs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoreAvailableVscsResponseBodyVscs struct {
	// The VSC device ID.
	//
	// example:
	//
	// vsc-xxxxx
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// The current status of the VSC device.
	//
	// example:
	//
	// Available
	VscStatus *string `json:"VscStatus,omitempty" xml:"VscStatus,omitempty"`
}

func (s ListKVCacheStoreAvailableVscsResponseBodyVscs) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoreAvailableVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoreAvailableVscsResponseBodyVscs) GetVscId() *string {
	return s.VscId
}

func (s *ListKVCacheStoreAvailableVscsResponseBodyVscs) GetVscStatus() *string {
	return s.VscStatus
}

func (s *ListKVCacheStoreAvailableVscsResponseBodyVscs) SetVscId(v string) *ListKVCacheStoreAvailableVscsResponseBodyVscs {
	s.VscId = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponseBodyVscs) SetVscStatus(v string) *ListKVCacheStoreAvailableVscsResponseBodyVscs {
	s.VscStatus = &v
	return s
}

func (s *ListKVCacheStoreAvailableVscsResponseBodyVscs) Validate() error {
	return dara.Validate(s)
}
