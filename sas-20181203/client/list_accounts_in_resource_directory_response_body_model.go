// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountsInResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccounts(v []*ListAccountsInResourceDirectoryResponseBodyAccounts) *ListAccountsInResourceDirectoryResponseBody
	GetAccounts() []*ListAccountsInResourceDirectoryResponseBodyAccounts
	SetRequestId(v string) *ListAccountsInResourceDirectoryResponseBody
	GetRequestId() *string
}

type ListAccountsInResourceDirectoryResponseBody struct {
	// La liste des comptes membres dans le répertoire de ressources.
	Accounts []*ListAccountsInResourceDirectoryResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
	// L\\"identifiant de la requête. Il s\\"agit d\\"un identifiant unique généré par Alibaba Cloud pour la requête. Vous pouvez l\\"utiliser pour le dépannage.
	//
	// example:
	//
	// B7A839E8-70AE-591D-8D9E-C5419A2240DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAccountsInResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsInResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountsInResourceDirectoryResponseBody) GetAccounts() []*ListAccountsInResourceDirectoryResponseBodyAccounts {
	return s.Accounts
}

func (s *ListAccountsInResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccountsInResourceDirectoryResponseBody) SetAccounts(v []*ListAccountsInResourceDirectoryResponseBodyAccounts) *ListAccountsInResourceDirectoryResponseBody {
	s.Accounts = v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBody) SetRequestId(v string) *ListAccountsInResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBody) Validate() error {
	if s.Accounts != nil {
		for _, item := range s.Accounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAccountsInResourceDirectoryResponseBodyAccounts struct {
	// L\\"identifiant de compte Alibaba Cloud du membre.
	//
	// example:
	//
	// 180803538814****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Deprecated
	//
	// Ce paramètre est obsolète et ne doit plus être utilisé.
	//
	// example:
	//
	// true.
	BuySas *int32 `json:"BuySas,omitempty" xml:"BuySas,omitempty"`
	// Indique si une instance Security Center en abonnement a été souscrite. Valeurs valides :
	//
	// - **true**
	//
	// - **false**.
	//
	// example:
	//
	// true.
	BuySasNew *bool `json:"BuySasNew,omitempty" xml:"BuySasNew,omitempty"`
	// Le mode de facturation de l\\"abonnement. Valeurs valides :
	//
	// 	- **0*	- : prépayé
	//
	// 	- **1*	- : paiement à l\\"usage.
	//
	// example:
	//
	// 1
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Le nom du compte.
	//
	// example:
	//
	// abc**
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// L\\"identifiant du dossier dans le répertoire de ressources.
	//
	// example:
	//
	// fd-cE2SQP****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// Le type d\\"achat de l\\"instance. Valeurs valides :
	//
	// - **0*	- : achat direct
	//
	// - **1*	- : alloué via la gestion multi-comptes.
	//
	// example:
	//
	// 0
	InstanceBuyType *int32 `json:"InstanceBuyType,omitempty" xml:"InstanceBuyType,omitempty"`
	// Indique si le compte est le compte de gestion du répertoire de ressources. Valeurs valides :
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsMaAccount *string `json:"IsMaAccount,omitempty" xml:"IsMaAccount,omitempty"`
	// Indique si le compte est marqué comme suivi.
	//
	// example:
	//
	// no
	IsMarked *string `json:"IsMarked,omitempty" xml:"IsMarked,omitempty"`
	// Indique si le compte est un compte administrateur délégué de Security Center. Valeurs valides :
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSasDaAccount *string `json:"IsSasDaAccount,omitempty" xml:"IsSasDaAccount,omitempty"`
	// Indique si le compte est un compte de gestion de Cloud Threat Detection and Response (CTDR). Valeurs valides :
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSiemControlAccount *string `json:"IsSiemControlAccount,omitempty" xml:"IsSiemControlAccount,omitempty"`
	// Indique si le compte est un compte administrateur délégué de Cloud Threat Detection and Response (CTDR). Valeurs valides :
	//
	// - **yes**
	//
	// - **no**.
	//
	// example:
	//
	// no
	IsSiemDaAccount *string `json:"IsSiemDaAccount,omitempty" xml:"IsSiemDaAccount,omitempty"`
	// Le commutateur de module à paiement à l\\"usage. Valeurs valides :
	//
	// - **0*	- : désactivé
	//
	// - **1*	- : activé.
	//
	// example:
	//
	// 1
	PostBasicService *int32 `json:"PostBasicService,omitempty" xml:"PostBasicService,omitempty"`
	// L\\"état des commutateurs de modules à paiement à l\\"usage, au format JsonString. Valeurs valides :
	//
	// - Clé :
	//
	//   - **VUL*	- : module de gestion des vulnérabilités
	//
	//   - **CSPM*	- : module Cloud Security Posture Management (CSPM)
	//
	//   - **AGENTLESS*	- : module de détection sans agent
	//
	//   - **SERVERLESS*	- : module de ressources serverless
	//
	//   - **CTDR*	- : module Cloud Threat Detection and Response (CTDR)
	//
	//   - **RASP*	- : module Runtime Application Self-Protection (RASP)
	//
	//   - **SDK*	- : module SDK de détection de fichiers malveillants
	//
	//   - **POST_HOST*	- : module de sécurité des hôtes et des conteneurs
	//
	// - Valeur : 0 indique désactivé. 1 indique activé.
	//
	// example:
	//
	// {\\"VUL\\":0}
	PostPayModuleSwitch *string `json:"PostPayModuleSwitch,omitempty" xml:"PostPayModuleSwitch,omitempty"`
	// L\\"identifiant de l\\"instance Security Center.
	//
	// example:
	//
	// postpay-sas-x5x3hy1ly***
	SaleInstance *string `json:"SaleInstance,omitempty" xml:"SaleInstance,omitempty"`
	// L\\"édition souscrite de Security Center. Valeurs valides :
	//
	// - **0*	- ou **1*	- : Édition gratuite
	//
	// - **2*	- ou **3*	- : Édition Enterprise
	//
	// - **5*	- : Édition Premium
	//
	// - **6*	- : Édition Anti-virus
	//
	// - **7*	- : Édition Ultimate.
	//
	// example:
	//
	// 0
	SasVersion *string `json:"SasVersion,omitempty" xml:"SasVersion,omitempty"`
}

func (s ListAccountsInResourceDirectoryResponseBodyAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListAccountsInResourceDirectoryResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetAccountId() *string {
	return s.AccountId
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetBuySas() *int32 {
	return s.BuySas
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetBuySasNew() *bool {
	return s.BuySasNew
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetChargeType() *int32 {
	return s.ChargeType
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetFolderId() *string {
	return s.FolderId
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetInstanceBuyType() *int32 {
	return s.InstanceBuyType
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsMaAccount() *string {
	return s.IsMaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsMarked() *string {
	return s.IsMarked
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSasDaAccount() *string {
	return s.IsSasDaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSiemControlAccount() *string {
	return s.IsSiemControlAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetIsSiemDaAccount() *string {
	return s.IsSiemDaAccount
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetPostBasicService() *int32 {
	return s.PostBasicService
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetPostPayModuleSwitch() *string {
	return s.PostPayModuleSwitch
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetSaleInstance() *string {
	return s.SaleInstance
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) GetSasVersion() *string {
	return s.SasVersion
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetAccountId(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.AccountId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetBuySas(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.BuySas = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetBuySasNew(v bool) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.BuySasNew = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetChargeType(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.ChargeType = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetDisplayName(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.DisplayName = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetFolderId(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.FolderId = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetInstanceBuyType(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.InstanceBuyType = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsMaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsMaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsMarked(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsMarked = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSasDaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSasDaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSiemControlAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSiemControlAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetIsSiemDaAccount(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.IsSiemDaAccount = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetPostBasicService(v int32) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.PostBasicService = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetPostPayModuleSwitch(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.PostPayModuleSwitch = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetSaleInstance(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.SaleInstance = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) SetSasVersion(v string) *ListAccountsInResourceDirectoryResponseBodyAccounts {
	s.SasVersion = &v
	return s
}

func (s *ListAccountsInResourceDirectoryResponseBodyAccounts) Validate() error {
	return dara.Validate(s)
}
