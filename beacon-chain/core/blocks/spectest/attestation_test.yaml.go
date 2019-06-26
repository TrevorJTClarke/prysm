// Code generated by yaml_to_go. DO NOT EDIT.
// source: attestation_minimal.yaml

package spectest

type AttestationTest struct {
	Title         string   `json:"title"`
	Summary       string   `json:"summary"`
	ForksTimeline string   `json:"forks_timeline"`
	Forks         []string `json:"forks"`
	Config        string   `json:"config"`
	Runner        string   `json:"runner"`
	Handler       string   `json:"handler"`
	TestCases     []struct {
		Description string `json:"description"`
		Pre         struct {
			Slot        uint64 `json:"slot"`
			GenesisTime uint64 `json:"genesis_time"`
			Fork        struct {
				PreviousVersion []byte `json:"previous_version"`
				CurrentVersion  []byte `json:"current_version"`
				Epoch           uint64 `json:"epoch"`
			} `json:"fork"`
			ValidatorRegistry []struct {
				Pubkey                     []byte `json:"pubkey" ssz:"size=48"`
				WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz:"size=32"`
				ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
				ActivationEpoch            uint64 `json:"activation_epoch"`
				ExitEpoch                  uint64 `json:"exit_epoch"`
				WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
				Slashed                    bool   `json:"slashed"`
				EffectiveBalance           uint64 `json:"effective_balance"`
			} `json:"validator_registry"`
			Balances                  []uint64      `json:"balances"`
			LatestRandaoMixes         [][]byte      `json:"latest_randao_mixes"`
			LatestStartShard          uint64        `json:"latest_start_shard"`
			PreviousEpochAttestations []interface{} `json:"previous_epoch_attestations"`
			CurrentEpochAttestations  []interface{} `json:"current_epoch_attestations"`
			PreviousJustifiedEpoch    uint64        `json:"previous_justified_epoch"`
			CurrentJustifiedEpoch     uint64        `json:"current_justified_epoch"`
			PreviousJustifiedRoot     []byte        `json:"previous_justified_root" ssz:"size=32"`
			CurrentJustifiedRoot      []byte        `json:"current_justified_root" ssz:"size=32"`
			JustificationBitfield     uint64        `json:"justification_bitfield"`
			FinalizedEpoch            uint64        `json:"finalized_epoch"`
			FinalizedRoot             []byte        `json:"finalized_root" ssz:"size=32"`
			CurrentCrosslinks         []struct {
				Shard      uint64 `json:"shard"`
				StartEpoch uint64 `json:"start_epoch"`
				EndEpoch   uint64 `json:"end_epoch"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				DataRoot   []byte `json:"data_root" ssz:"size=32"`
			} `json:"current_crosslinks"`
			PreviousCrosslinks []struct {
				Shard      uint64 `json:"shard"`
				StartEpoch uint64 `json:"start_epoch"`
				EndEpoch   uint64 `json:"end_epoch"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				DataRoot   []byte `json:"data_root" ssz:"size=32"`
			} `json:"previous_crosslinks"`
			LatestBlockRoots       [][]byte `json:"latest_block_roots"`
			LatestStateRoots       [][]byte `json:"latest_state_roots"`
			LatestActiveIndexRoots [][]byte `json:"latest_active_index_roots"`
			LatestSlashedBalances  []uint64 `json:"latest_slashed_balances"`
			LatestBlockHeader      struct {
				Slot       uint64 `json:"slot"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				StateRoot  []byte `json:"state_root" ssz:"size=32"`
				BodyRoot   []byte `json:"body_root" ssz:"size=32"`
				Signature  []byte `json:"signature" ssz:"size=96"`
			} `json:"latest_block_header"`
			HistoricalRoots []interface{} `json:"historical_roots"`
			LatestEth1Data  struct {
				DepositRoot  []byte `json:"deposit_root" ssz:"size=32"`
				DepositCount uint64 `json:"deposit_count"`
				BlockHash    []byte `json:"block_hash" ssz:"size=32"`
			} `json:"latest_eth1_data"`
			Eth1DataVotes []struct {
				DepositRoot  []byte `json:"deposit_root" ssz:"size=32"`
				DepositCount uint64 `json:"deposit_count"`
				BlockHash    []byte `json:"block_hash" ssz:"size=32"`
			} `json:"eth1_data_votes"`
			DepositIndex uint64 `json:"deposit_index"`
		} `json:"pre"`
		Attestation struct {
			AggregationBitfield []byte `json:"aggregation_bitfield"`
			Data                struct {
				BeaconBlockRoot []byte `json:"beacon_block_root" ssz:"size=32"`
				SourceEpoch     uint64 `json:"source_epoch"`
				SourceRoot      []byte `json:"source_root" ssz:"size=32"`
				TargetEpoch     uint64 `json:"target_epoch"`
				TargetRoot      []byte `json:"target_root" ssz:"size=32"`
				Crosslink       struct {
					Shard      uint64 `json:"shard"`
					StartEpoch uint64 `json:"start_epoch"`
					EndEpoch   uint64 `json:"end_epoch"`
					ParentRoot []byte `json:"parent_root" ssz:"size=32"`
					DataRoot   []byte `json:"data_root" ssz:"size=32"`
				} `json:"crosslink"`
			} `json:"data"`
			CustodyBitfield []byte `json:"custody_bitfield"`
			Signature       []byte `json:"signature" ssz:"size=96"`
		} `json:"attestation"`
		Post struct {
			Slot        uint64 `json:"slot"`
			GenesisTime uint64 `json:"genesis_time"`
			Fork        struct {
				PreviousVersion []byte `json:"previous_version"`
				CurrentVersion  []byte `json:"current_version"`
				Epoch           uint64 `json:"epoch"`
			} `json:"fork"`
			ValidatorRegistry []struct {
				Pubkey                     []byte `json:"pubkey" ssz:"size=48"`
				WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz:"size=32"`
				ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
				ActivationEpoch            uint64 `json:"activation_epoch"`
				ExitEpoch                  uint64 `json:"exit_epoch"`
				WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
				Slashed                    bool   `json:"slashed"`
				EffectiveBalance           uint64 `json:"effective_balance"`
			} `json:"validator_registry"`
			Balances                  []uint64      `json:"balances"`
			LatestRandaoMixes         [][]byte      `json:"latest_randao_mixes"`
			LatestStartShard          uint64        `json:"latest_start_shard"`
			PreviousEpochAttestations []interface{} `json:"previous_epoch_attestations"`
			CurrentEpochAttestations  []interface{} `json:"current_epoch_attestations"`
			PreviousJustifiedEpoch    uint64        `json:"previous_justified_epoch"`
			CurrentJustifiedEpoch     uint64        `json:"current_justified_epoch"`
			PreviousJustifiedRoot     []byte        `json:"previous_justified_root" ssz:"size=32"`
			CurrentJustifiedRoot      []byte        `json:"current_justified_root" ssz:"size=32"`
			JustificationBitfield     uint64        `json:"justification_bitfield"`
			FinalizedEpoch            uint64        `json:"finalized_epoch"`
			FinalizedRoot             []byte        `json:"finalized_root" ssz:"size=32"`
			CurrentCrosslinks         []struct {
				Shard      uint64 `json:"shard"`
				StartEpoch uint64 `json:"start_epoch"`
				EndEpoch   uint64 `json:"end_epoch"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				DataRoot   []byte `json:"data_root" ssz:"size=32"`
			} `json:"current_crosslinks"`
			PreviousCrosslinks []struct {
				Shard      uint64 `json:"shard"`
				StartEpoch uint64 `json:"start_epoch"`
				EndEpoch   uint64 `json:"end_epoch"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				DataRoot   []byte `json:"data_root" ssz:"size=32"`
			} `json:"previous_crosslinks"`
			LatestBlockRoots       [][]byte `json:"latest_block_roots"`
			LatestStateRoots       [][]byte `json:"latest_state_roots"`
			LatestActiveIndexRoots [][]byte `json:"latest_active_index_roots"`
			LatestSlashedBalances  []uint64 `json:"latest_slashed_balances"`
			LatestBlockHeader      struct {
				Slot       uint64 `json:"slot"`
				ParentRoot []byte `json:"parent_root" ssz:"size=32"`
				StateRoot  []byte `json:"state_root" ssz:"size=32"`
				BodyRoot   []byte `json:"body_root" ssz:"size=32"`
				Signature  []byte `json:"signature" ssz:"size=96"`
			} `json:"latest_block_header"`
			HistoricalRoots []interface{} `json:"historical_roots"`
			LatestEth1Data  struct {
				DepositRoot  []byte `json:"deposit_root" ssz:"size=32"`
				DepositCount uint64 `json:"deposit_count"`
				BlockHash    []byte `json:"block_hash" ssz:"size=32"`
			} `json:"latest_eth1_data"`
			Eth1DataVotes []struct {
				DepositRoot  []byte `json:"deposit_root" ssz:"size=32"`
				DepositCount uint64 `json:"deposit_count"`
				BlockHash    []byte `json:"block_hash" ssz:"size=32"`
			} `json:"eth1_data_votes"`
			DepositIndex uint64 `json:"deposit_index"`
		} `json:"post"`
		BlsSetting uint64 `json:"bls_setting,omitempty"`
	} `json:"test_cases"`
}
