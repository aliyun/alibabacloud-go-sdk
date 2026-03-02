// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcRepoForkListResult interface {
	dara.Model
	String() string
	GoString() string
	SetReposForks(v []*ReposFork) *PbcRepoForkListResult
	GetReposForks() []*ReposFork
	SetRequestId(v string) *PbcRepoForkListResult
	GetRequestId() *string
}

type PbcRepoForkListResult struct {
	ReposForks []*ReposFork `json:"reposForks,omitempty" xml:"reposForks,omitempty" type:"Repeated"`
	RequestId  *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcRepoForkListResult) String() string {
	return dara.Prettify(s)
}

func (s PbcRepoForkListResult) GoString() string {
	return s.String()
}

func (s *PbcRepoForkListResult) GetReposForks() []*ReposFork {
	return s.ReposForks
}

func (s *PbcRepoForkListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcRepoForkListResult) SetReposForks(v []*ReposFork) *PbcRepoForkListResult {
	s.ReposForks = v
	return s
}

func (s *PbcRepoForkListResult) SetRequestId(v string) *PbcRepoForkListResult {
	s.RequestId = &v
	return s
}

func (s *PbcRepoForkListResult) Validate() error {
	if s.ReposForks != nil {
		for _, item := range s.ReposForks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
