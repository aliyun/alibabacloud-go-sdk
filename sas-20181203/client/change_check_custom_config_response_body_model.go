// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckCustomConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIllegalCustomConfigs(v []*ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) *ChangeCheckCustomConfigResponseBody
	GetIllegalCustomConfigs() []*ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs
	SetIllegalRepairConfigs(v []*ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) *ChangeCheckCustomConfigResponseBody
	GetIllegalRepairConfigs() []*ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs
	SetRequestId(v string) *ChangeCheckCustomConfigResponseBody
	GetRequestId() *string
}

type ChangeCheckCustomConfigResponseBody struct {
	// An array that consists of the invalid custom configuration items of the check item.
	IllegalCustomConfigs []*ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs `json:"IllegalCustomConfigs,omitempty" xml:"IllegalCustomConfigs,omitempty" type:"Repeated"`
	// An array that consists of the invalid parameters required for fixing risk items.
	IllegalRepairConfigs []*ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs `json:"IllegalRepairConfigs,omitempty" xml:"IllegalRepairConfigs,omitempty" type:"Repeated"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// DA8133CC-CCA0-5CF2-BF64-FE7D52C44***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeCheckCustomConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigResponseBody) GetIllegalCustomConfigs() []*ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs {
	return s.IllegalCustomConfigs
}

func (s *ChangeCheckCustomConfigResponseBody) GetIllegalRepairConfigs() []*ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs {
	return s.IllegalRepairConfigs
}

func (s *ChangeCheckCustomConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCheckCustomConfigResponseBody) SetIllegalCustomConfigs(v []*ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) *ChangeCheckCustomConfigResponseBody {
	s.IllegalCustomConfigs = v
	return s
}

func (s *ChangeCheckCustomConfigResponseBody) SetIllegalRepairConfigs(v []*ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) *ChangeCheckCustomConfigResponseBody {
	s.IllegalRepairConfigs = v
	return s
}

func (s *ChangeCheckCustomConfigResponseBody) SetRequestId(v string) *ChangeCheckCustomConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCheckCustomConfigResponseBody) Validate() error {
	if s.IllegalCustomConfigs != nil {
		for _, item := range s.IllegalCustomConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.IllegalRepairConfigs != nil {
		for _, item := range s.IllegalRepairConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs struct {
	// The name of the custom configuration item, which is unique in a check item.
	//
	// example:
	//
	// SessionTimeMax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) GetName() *string {
	return s.Name
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) SetName(v string) *ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs {
	s.Name = &v
	return s
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalCustomConfigs) Validate() error {
	return dara.Validate(s)
}

type ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs struct {
	// The name of the invalid parameter required for fixing a risk item.
	//
	// example:
	//
	// SessionTimeMax
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) GoString() string {
	return s.String()
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) GetName() *string {
	return s.Name
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) SetName(v string) *ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs {
	s.Name = &v
	return s
}

func (s *ChangeCheckCustomConfigResponseBodyIllegalRepairConfigs) Validate() error {
	return dara.Validate(s)
}
